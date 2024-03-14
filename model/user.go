package forum

import (
	"database/sql"
	"errors"
)

// TABLE: users

// Insert new row into user table
func UserInsertDb(applicant Applicant, db *sql.DB, pass []byte) error {
	sqlStmt, err := db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer sqlStmt.Close()
	_, err = sqlStmt.Exec(applicant.Username, applicant.Email, pass, "member")
	if err != nil {
		return err
	}
	return nil
}

// Check if email exists in user table
func UserExistsDb(applicantEmail string) error {
	sqlStmt := `SELECT EXISTS (SELECT 1 FROM users WHERE user_email = ?)`
	var exists bool
	err := DB.QueryRow(sqlStmt, applicantEmail).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return UserExistsError
	}
	return nil
}

// Check if username exists in user table
func UsernameExistsDb(applicantUsername string) error {
	sqlStmt := `SELECT EXISTS (SELECT 1 FROM users WHERE user_name = ?)`
	var exists bool
	err := DB.QueryRow(sqlStmt, applicantUsername).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return UsernameExistsError
	}
	return nil
}

// Retrieve data from the user table and assign into the global struct
func UserRetrieveDb(email string, password string) error {
	userdata := DB.QueryRow("SELECT user_id, user_name, user_pass, user_email, user_type FROM users where user_email = ?", email) // select gets the data from users table
	err := userdata.Scan(&LoggedUser.Userid, &LoggedUser.Username, &LoggedUser.Password, &LoggedUser.Email, &LoggedUser.Type)     // scan assigns the data of the row to variables
	if err != nil {
		return err
	}
	return nil
}

// get post username from database
func GetPostUsername(p *Post) error {
	userData := DB.QueryRow("Select user_name from users where user_id = ?", p.UserID)
	err := userData.Scan(&p.Username)
	if err != nil {
		return errors.New("User Scan Error:" + err.Error())
	}
	return nil
}

// get comment username from database
func GetCommentUsername(c *Comment) error {
	userData := DB.QueryRow("Select user_name from users where user_id = ?", c.User_id)
	err := userData.Scan(&c.CommentUsername)
	if err != nil {
		return errors.New("Comment Username Scan Error:" + err.Error())
	}
	return nil
}
