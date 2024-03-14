package forum

import (
	"database/sql"
	"net/http"

	bcrypt "golang.org/x/crypto/bcrypt"
)

// Receive new user data, validate and insert in user table
func UserRegisteration(applicant Applicant, db *sql.DB) error {
	err := RegisterValidator(applicant)
	if err != nil {
		AllData.LoginErrorMsg = err.Error()
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(applicant.Password), 4)
	if err != nil {
		AllData.LoginErrorMsg = UserPasswordError.Error() // Using simplified error for user
		return err
	}
	err = UserInsertDb(applicant, db, pass)
	if err != nil {
		return err
	}
	return nil
}

// Receive login credentials, validate and respond with a session cookie
func UserLogin(email string, password string) (*http.Cookie, error) {
	// Validate User Existance
	if UserExistsDb(email) != nil {
		AllData.LoginErrorMsg = UserEmailError.Error()
		return BlankCookie, UserEmailError
	}
	// Retrieves User Data
	err := UserRetrieveDb(email, password)
	if err != nil {
		return BlankCookie, err
	}
	// Validates Entered Password
	err = bcrypt.CompareHashAndPassword([]byte(LoggedUser.Password), []byte(password))
	if err != nil {
		AllData.LoginErrorMsg = UserPasswordError.Error()
		return BlankCookie, err
	}

	LoggedUser.Registered = true
	AllData.IsLogged = true
	AllData.LoggedUser = LoggedUser
	AllData.LoggedUserID = LoggedUser.Userid

	if LoggedUser.Type == "admin" {
		AllData.TypeAdmin = true
	}
	cookie, err := CreateSession()
	if err != nil {
		return BlankCookie, err
	}
	ActiveUsersData[LoggedUser.Userid] = AllData
	err = GetUserPostsInteractions()
	if err != nil {
		return BlankCookie, err
	}
	return cookie, nil
}
