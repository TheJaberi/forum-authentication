package forum

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

var (
	Database      *sql.DB
	LoggedUser    User
	ErrorMsg      string
	AllCategories []Category
	AllData       Data
	AllPosts      []Post
	LiveSession   Session
	EmptySession  Session
	Empty         User
	LoginError2   bool
	PostError2    bool
)

type Data struct {
	AllPosts      []Post
	AllCategories []Category
	Postpage      Post
	LoggedUser    User
	CategoryCheck bool
	IsLogged      bool
	LoggedUserID  int
	TypeAdmin     bool
	LoginError    bool
	LoginErrorMsg string
	PostError     bool
	PostErrorMsg  string
	Port          int
}
type Category struct {
	CategoryName string
	CategoryID   int
}
type User struct {
	Userid     int
	Username   string
	Password   string
	Email      string
	Registered bool
	Type       string
}

var ErrResponse struct {
	StatusCode bool
	ErrorMsg   string
}

type Post struct {
	Title          string
	Body           string
	PostID         int
	UserID         int
	Username       string
	Category       []Category
	Likes          int
	Dislikes       int
	Userlike       bool
	UserDislike    bool
	LoggedUser     bool
	TimeCreated    string
	Comments       []Comment
	NumbOfComments int
	Image          string
	Port           int
}

type Comment struct {
	Body               string
	Post_id            int
	User_id            int
	CommentUsername    string
	TimeCreated        string
	Likes              int
	Dislikes           int
	CommentUserlike    bool
	CommentUserDislike bool
	CommentLoggedUser  bool
	Comment_id         int
}

type Applicant struct {
	Username string
	Email    string
	Password []byte
	Type     string
}

type Session struct {
	User      string
	UserId    int
	Uuid      uuid.UUID
	CreatedAt time.Time
	Expires   time.Time
}

var ActiveSessions map[string]Session = make(map[string]Session)

var BlankCookie = &http.Cookie{
	Name:     "session_token",
	Value:    "",
	Domain:   "localhost",
	Path:     "/",
	MaxAge:   -1,
	HttpOnly: true,
}

var ActiveUsersData map[int]Data = make(map[int]Data)

var SessionLinks map[int]uuid.UUID = make(map[int]uuid.UUID)
