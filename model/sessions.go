package forum

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func CreateSession() (*http.Cookie, error) {
	// Create New Token Value
	uuid, err := uuid.NewV4()
	if err != nil {
		return BlankCookie, err
	}
	_, exists := ActiveUsersData[LoggedUser.Userid]
	if exists {
		delete(ActiveUsersData, ActiveSessions[SessionLinks[LoggedUser.Userid].String()].UserId)
		delete(ActiveSessions, SessionLinks[LoggedUser.Userid].String())
		delete(SessionLinks, LoggedUser.Userid)
	}

	// Add Session Token Details To Active Sessions Map
	ActiveSessions[uuid.String()] = Session{
		User:      LoggedUser.Username,
		UserId:    LoggedUser.Userid,
		Uuid:      uuid,
		CreatedAt: time.Now(),
		Expires:   time.Now().Add(3600 * time.Second),
	}

	SessionLinks[LoggedUser.Userid] = uuid

	// Assign Token Value in Cookie
	c := &http.Cookie{
		Name:     "session_token",
		Value:    uuid.String(),
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	return c, nil
}

func (s Session) isExpired() bool {
	return s.Expires.Before(time.Now())
}

// Validate session remaining time
func ValidateSession(r *http.Request) error {
	c, err := r.Cookie("session_token")
	if err != nil {
		return err
	}
	sessionToken := c.Value
	userSession, exists := ActiveSessions[sessionToken]
	if !exists && AllData.IsLogged {
		// If the session token is not present in session map, return an unauthorized error
		return SessionInvalid
	}
	if userSession.isExpired() || c.MaxAge < 0 {
		RemoveSession(r)
		return err
	}
	if LoadSession(userSession.UserId) != nil {
		RemoveSession(r)
		return err
	}
	return nil
}

func LoadSession(userID int) error {
	data, exists := ActiveUsersData[userID]
	if !exists {
		return ActiveUserError
	}
	AllData.LoggedUser = data.LoggedUser
	AllData.IsLogged = data.IsLogged
	AllData.Postpage.UserID = data.Postpage.UserID
	AllData.LoggedUserID = data.LoggedUserID

	// GetUserPostsInteractions()
	err := GetPosts()
	if err != nil {
		return err
	}
	GetUserCommentInteractions()
	return nil
}

func RemoveSession(r *http.Request) int {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// log.Print()
			AllData.LoggedUser = Empty
			AllData.IsLogged = false
			AllData.TypeAdmin = false
			LiveSession = EmptySession
			return 0
		}
		// For any other type of error, return a bad request status
		return http.StatusBadRequest
	}
	sessionToken := c.Value
	delete(ActiveUsersData, ActiveSessions[sessionToken].UserId)
	delete(ActiveSessions, sessionToken)
	AllData.LoggedUser = Empty
	AllData.IsLogged = false
	AllData.TypeAdmin = false
	LiveSession = EmptySession
	return 0
}
