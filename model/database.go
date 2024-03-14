package forum

import (
	"database/sql"
	"log"
	"os"
)

// creates all the tables in the database
func CreateDBTables(db *sql.DB) {
	UsersTable(db)
	PostsTable(db)
	PostInteractionsTable(db)
	CommentsTable(db)
	CommentInteractionsTable(db)
	RequestsTable(db)
	ActionsTable(db)
	CategoryTable(db)
	PostCategoryTable(db)
}

func UsersTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER NOT NULL,
		user_name CHAR(10) NOT NULL UNIQUE,
		user_email CHAR(25) NOT NULL UNIQUE,
		user_pass PASSWORD NOT NULL,
		user_type TEXT NOT NULL DEFAULT member,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("user_id" AUTOINCREMENT)
	);`
	usersTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		os.Exit(0)
	}
	defer usersTable.Close()
	_, err = usersTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func PostsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		Title TEXT NOT NULL,
		img_url TEXT,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	);`
	postsTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer postsTable.Close()
	_, err = postsTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func PostInteractionsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS Interaction (
		id INTEGER PRIMARY KEY,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		interaction BIT NOT NULL);`
	interaction_postsTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer interaction_postsTable.Close()
	_, err = interaction_postsTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func CommentsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS comments (
		comment_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("comment_id" AUTOINCREMENT),
		FOREIGN KEY("post_id") REFERENCES posts("post_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	);`
	commentsTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer commentsTable.Close()
	_, err = commentsTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func CommentInteractionsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS interaction_comments (
		comment_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		interaction TEXT NOT NULL,
		FOREIGN KEY("comment_id") REFERENCES comments("comment_id"),
		FOREIGN KEY("post_id") REFERENCES posts("post_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	interaction_commentsTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer interaction_commentsTable.Close()
	_, err = interaction_commentsTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func RequestsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS requests (
		request_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY("request_id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	requestsTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer requestsTable.Close()
	_, err = requestsTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func ActionsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS actions (
		request_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		body TEXT NOT NULL,
		time_created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY("request_id") REFERENCES requests("request_id"),
		FOREIGN KEY("user_id") REFERENCES users("user_id")
	)`
	actionsTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer actionsTable.Close()
	_, err = actionsTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func CategoryTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS Category (
		id INTEGER PRIMARY KEY,
		Name TEXT
	)`
	categoryTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer categoryTable.Close()
	_, err = categoryTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
func PostCategoryTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS Post2Category
	(id INTEGER PRIMARY KEY,
	post_id INTEGER NOT NULL,
	category_id INTEGER NOT NULL
)`
	post2categoryTable, err := db.Prepare(query)
	if err != nil {
		log.Println(err.Error())
	}
	defer post2categoryTable.Close()
	_, err = post2categoryTable.Exec()
	if err != nil {
		log.Println(err.Error())
	}
}
