package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

// loads the database and checks if the tables are created
func DatabaseLoader() {
	var err error
	DB, err = sql.Open("sqlite3", "../model/forum.db")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if !TableExists(DB, "users") {
		CreateDBTables(DB)
		sqlStmt, err := DB.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer sqlStmt.Close()
		admin_pass, err := bcrypt.GenerateFromPassword([]byte("adminpass"), 4)
		if err != nil {
			log.Println(err.Error())
		}
		sqlStmt.Exec("admin", "admin@gmail.com", admin_pass, "admin")
	}

}

// checks if the tables exist
func TableExists(db *sql.DB, tableName string) bool {
	sqlStmt, err := db.Prepare("SELECT COUNT(*) FROM sqlite_master WHERE type = 'table' AND name = ?")
	if err != nil {
		return false
	}
	defer sqlStmt.Close()
	var count int
	err = sqlStmt.QueryRow(tableName).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
