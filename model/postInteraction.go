package forum

import (
	"context"
	"errors"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// post interaction handles the like and dislike buttons
func PostInteractions(add, remove, path string) (Post, error) {
	var p Post
	var addPost_id int
	var remPost_id int
	var err error
	if add != "" {
		addPost_id, err = strconv.Atoi(add) // post interaction handles the data from like or dislike button if the user logged hasn't already clicked on it
		if err != nil {
			log.Println(err)
			return p, err
		}

	}
	if remove != "" {
		remPost_id, err = strconv.Atoi(remove) // remove interaction handles the data from like or dislike button if the user logged has already clicked on it
		if err != nil {
			log.Println(err)
			return p, err
		}
	}
	user_id := AllData.LoggedUser.Userid
	if addPost_id > remPost_id { // which ever value is greater determines whether to add or remove
		if path == "/like/" {
			if !AllPosts[addPost_id-1].UserDislike {
				InsertPostInteraction(addPost_id, user_id, 1) // insert adds the interaction to the database 1 is like 0 is dislike
				AllPosts[addPost_id-1].Userlike = true        // changes the post like or dislike for the logged in user in the all posts var
				AllPosts[addPost_id-1].Likes++
			} else {
				UpdatePostInteraction(addPost_id, user_id, 1) // update is used if a like has to be changed to a dislike or vice versa
				AllPosts[addPost_id-1].Userlike = true
				AllPosts[addPost_id-1].UserDislike = false
				AllPosts[addPost_id-1].Likes++
				AllPosts[addPost_id-1].Dislikes--
			}
		} else {
			if !AllPosts[addPost_id-1].Userlike {
				InsertPostInteraction(addPost_id, user_id, 0)
				AllPosts[addPost_id-1].UserDislike = true
				AllPosts[addPost_id-1].Dislikes++
			} else {
				UpdatePostInteraction(addPost_id, user_id, 0)
				AllPosts[addPost_id-1].UserDislike = true
				AllPosts[addPost_id-1].Userlike = false
				AllPosts[addPost_id-1].Likes--
				AllPosts[addPost_id-1].Dislikes++
			}
		}
		p = AllPosts[addPost_id-1]
	} else {
		RemovePostInteraction(remPost_id, user_id) //remove is greater means there is already an interaction that needs to be removed
		if AllPosts[remPost_id-1].Userlike {
			AllPosts[remPost_id-1].Userlike = false
			AllPosts[remPost_id-1].Likes--
		}
		if AllPosts[remPost_id-1].UserDislike {
			AllPosts[remPost_id-1].UserDislike = false
			AllPosts[remPost_id-1].Dislikes--
		}
		p = AllPosts[remPost_id-1]
	}
	p.LoggedUser = true
	return p, nil
}

// insert the interaction into the database
func InsertPostInteraction(postID int, userID int, likeOrDislike int) {
	query := "INSERT INTO `Interaction` (`post_id`, `user_id`, `interaction`) VALUES (?, ?, ?)"
	_, err := DB.ExecContext(context.Background(), query, postID, userID, likeOrDislike)
	if err != nil {
		log.Println(err)
	}
}

// removes the interaction into the database
func RemovePostInteraction(postID int, userID int) {
	query := "DELETE FROM `Interaction` where post_id = ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, postID, userID)
	if err != nil {
		log.Println(err)
	}
}

// updates the interaction in the data
func UpdatePostInteraction(postID int, userID int, likeOrDislike int) {
	query := "UPDATE Interaction SET interaction = ? where post_id= ? AND user_id = ?"
	_, err := DB.ExecContext(context.Background(), query, likeOrDislike, postID, userID)
	if err != nil {
		log.Println(err)
	}
}

// gets the users interaction from the database once he is logged in and adds it to the database
func GetUserPostsInteractions() error {
	for i := range AllPosts {
		var interaction int
		postData := DB.QueryRow("SELECT interaction from Interaction where post_id = ? AND user_id = ?", i+1, AllData.LoggedUser.Userid)
		err := postData.Scan(&interaction)
		if err != nil {
			continue // used for logout (remove user post interactions from global struct)?
		} else {
			if interaction == 1 {
				AllPosts[i].Userlike = true
			} else {
				AllPosts[i].UserDislike = true
			}
		}
	}
	return nil
}

// counts the likes for the post from the database
func GetPostLikes(p *Post) error {
	likedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", p.PostID, 1)
	err := likedata.Scan(&p.Likes)
	if err != nil {
		return errors.New("Post Interaction (Likes) Scan Error:" + err.Error())
	}
	return nil
}

// counts the dislikes for the post from the database
func GetPostDislikes(p *Post) error {
	dislikedata := DB.QueryRow("SELECT COUNT(user_id) FROM Interaction where post_id = ? AND interaction = ?", p.PostID, 0)
	err := dislikedata.Scan(&p.Dislikes)
	if err != nil {
		return errors.New("Post Interaction (Dislikes) Scan Error:" + err.Error())
	}
	return nil
}
