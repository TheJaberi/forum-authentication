package forum

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// create comment adds the comment to the database and also appends it to the post struct
func CreateComment(rawComment, postStrID string) (Post, error) {
	var postID int
	var err error
	rawComment = RemoveSpaces(rawComment)
	if rawComment == "" {
		return Post{}, errors.New("comment invalid")
	}
	if postStrID != "" {
		postID, err = strconv.Atoi(postStrID)
		if err != nil {
			return Post{}, err
		}
	}
	var c = Comment{
		Post_id:         postID,
		User_id:         AllData.LoggedUser.Userid,
		CommentUsername: AllData.LoggedUser.Username,
		Body:            rawComment,
	}
	id, err := CreateCommentDb(c)
	if err != nil {
		return Post{}, err
	}
	c, err = GetComment(id)
	if err != nil {
		log.Println(err.Error())
	}
	// Using AllPosts, otherswise AllData is sorted in reverese
	AllPosts[c.Post_id-1].Comments = append(AllPosts[c.Post_id-1].Comments, c)
	return AllPosts[c.Post_id-1], nil
}

func CreateCommentDb(c Comment) (int, error) {
	query := "INSERT INTO `comments` (`post_id`, `user_id`, `body`) VALUES (?, ?, ?)"
	rowData, err := DB.ExecContext(context.Background(), query, c.Post_id, c.User_id, c.Body)
	if err != nil {
		log.Println(err)
	}
	commentID, err := rowData.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(commentID), nil
}

// gets the comment using the ID
func GetComment(id int) (Comment, error) {
	row := DB.QueryRow("SELECT comment_id, post_id, user_id, body, time_created from comments WHERE comment_id=?", id)
	var c Comment
	err := row.Scan(&c.Comment_id, &c.Post_id, &c.User_id, &c.Body, &c.TimeCreated)
	if err != nil {
		log.Printf("Error Getting Post")
		return c, err
	}
	c.TimeCreated = strings.Replace(c.TimeCreated, "T", " ", -1)
	c.TimeCreated = strings.Replace(c.TimeCreated, "Z", " ", -1)
	err = GetCommentUsername(&c)
	if err != nil {
		log.Println(err.Error())
		return c, err
	}
	return c, nil
}

// get post comments adds the comments to each post
func GetPostComments(p *Post) error {
	p.Comments = nil
	commentData, err := DB.Query("Select comment_id, body, user_id, time_created from comments where post_id = ?", p.PostID)
	if err != nil {
		return errors.New("Comment Query Error:" + err.Error())
	}
	defer commentData.Close()
	for commentData.Next() {
		var c Comment
		err := commentData.Scan(&c.Comment_id, &c.Body, &c.User_id, &c.TimeCreated)
		if err != nil {
			return errors.New("Comment Scan Error:" + err.Error())
		}
		c.TimeCreated = strings.Replace(c.TimeCreated, "T", " ", -1)
		c.TimeCreated = strings.Replace(c.TimeCreated, "Z", " ", -1)
		err = GetCommentUsername(&c)
		if err != nil {
			return err
		}
		p.Comments = append(p.Comments, c)
	}
	p.NumbOfComments = len(p.Comments)
	return nil
}
