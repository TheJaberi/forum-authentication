package forum

import (
	"context"
	"errors"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// adds the post created and its categories to the database
func CreatePost(title string, body string, postCategories []int) error {
	title = RemoveSpaces(title)
	body = RemoveSpaces(body)
	if len(title) == 0 || len(body) == 0 {
		return PostError
	}
	var postData = Post{
		Title:    title,
		Body:     body,
		UserID:   LoggedUser.Userid,
		Username: LoggedUser.Username,
	}
	id := CreatePostDb(postData)
	err := AssignPostCategoryDb(id, postCategories)
	if err != nil {
		return err
	}
	newPost, err := GetPost(id)
	if err != nil {
		return err
	}
	AllData.AllPosts = append(AllData.AllPosts, newPost)
	return nil
}

// adds the post created to the database
func CreatePostDb(post Post) int {
	query := "INSERT INTO `posts` (`Title`, `body`, `user_id`) VALUES (?, ?, ?)"
	rowData, err := DB.ExecContext(context.Background(), query, post.Title, post.Body, post.UserID)
	if err != nil { // the post is added using the ExecContext along with the userid which is in the LoggedUser variable
		log.Println(err)
	}
	postID, err := rowData.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return int(postID)
}

// adds all the data for the post to a variable and returns it
func GetPost(id int) (Post, error) {
	row := DB.QueryRow("SELECT id, Title, body, user_id, time_created from posts WHERE id=?", id)
	var p Post
	err := row.Scan(&p.PostID, &p.Title, &p.Body, &p.UserID, &p.TimeCreated)
	if err != nil {
		log.Printf("Error Getting Post")
		return p, err
	}
	p.TimeCreated = strings.Replace(p.TimeCreated, "T", " ", -1)
	p.TimeCreated = strings.Replace(p.TimeCreated, "Z", " ", -1)
	p, err = GetPostDetails(p)
	if err != nil {
		log.Println(errors.New("Post Scan Error: " + err.Error()))
		return p, err
	}
	if LoggedUser.Registered {
		err := GetUserPostsInteractions()
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

// adds all the data for the post to a struct then appends them to an array
func GetPosts() error {
	AllPosts = nil
	postData, err := DB.Query("Select id, Title, body, user_id, time_created from posts")
	if err != nil {
		log.Println(err)
	}
	defer postData.Close()
	for postData.Next() {
		var p Post
		err := postData.Scan(&p.PostID, &p.Title, &p.Body, &p.UserID, &p.TimeCreated)
		if err != nil {
			log.Println(errors.New("Post Scan Error: " + err.Error()))
			return err
		}
		p.TimeCreated = strings.Replace(p.TimeCreated, "T", " ", -1)
		p.TimeCreated = strings.Replace(p.TimeCreated, "Z", " ", -1)
		p, err = GetPostDetails(p)
		if err != nil {
			log.Println(err.Error())
		}
		AllPosts = append(AllPosts, p)
	}
	if LoggedUser.Registered {
		err := GetUserPostsInteractions()
		if err != nil {
			return err
		}
	}
	AllData.AllPosts = AllPosts
	return nil
}

// gets the data required to fill the struct from every table
func GetPostDetails(p Post) (Post, error) {
	// USERNAME
	err := GetPostUsername(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// CATEGORIES
	err = GetPostCategories(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// COMMENTS
	err = GetPostComments(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	// INTERACTIONS
	err = GetPostLikes(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	err = GetPostDislikes(&p)
	if err != nil {
		log.Println(err.Error())
		return p, err
	}
	return p, nil
}
