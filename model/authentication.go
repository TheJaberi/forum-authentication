package forum

import (
	"database/sql"
	"net/http"

	bcrypt "golang.org/x/crypto/bcrypt"
)

// Receive new user data, validate and insert in user table
func UserRegisteration(applicant Applicant, db *sql.DB) error {
	err := RegisterValidator(applicant)
	if err == UsernameExistsError {
		if applicant.Reg_type == 1 {
			applicant.Username = applicant.Username + "(google)"
		} else if applicant.Reg_type == 2 {
			applicant.Username = applicant.Username + "(github)"
		}
	} else if err != nil {
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
func UserLogin(email string, password string, RegType int) (*http.Cookie, error) {
	// Validate User Existance
	if UserExistsDb(email, RegType) != nil {
		AllData.LoginErrorMsg = UserExistsError.Error()
		return BlankCookie, UserExistsError
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

func UserLoginGoogleAuth(email string, password string, RegType int) (*http.Cookie, error) {
	// Validate User Existance
	if UserExistsDb(email, RegType) == UserExistsError {
		AllData.LoginErrorMsg = UserExistsError.Error()
		return BlankCookie, UserEmailError
	} else if UserExistsDb(email, RegType) != nil {
		appicant := Applicant{Username:RemoveEmailDetails(email), Email: email, Password: []byte(password), Reg_type: RegType}
		err := UserRegisteration(appicant, DB)
		if err !=nil{
			AllData.LoginErrorMsg = err.Error()
		}
	}
	cookies, err := UserLogin(email, password, RegType)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}

func UserLoginGithubAuth(username string, email string, password string, RegType int) (*http.Cookie, error) {
	// Validate User Existance
	if UserExistsDb(email, RegType) == UserExistsError {
		AllData.LoginErrorMsg = UserExistsError.Error()
		return BlankCookie, UserEmailError
	} else if UserExistsDb(email, RegType) != nil {
		appicant := Applicant{Username: username, Email: email, Password: []byte(password), Reg_type: RegType}
		UserRegisteration(appicant, DB)
	}
	cookies, err := UserLogin(email, password, RegType)
	if err != nil {
		return nil, err
	}
	return cookies, nil
}

func RemoveEmailDetails(email string)string {
	emailByte := []byte(email)
	var finalUsername string
	for i:=0;i<len(emailByte);i++{
		if (emailByte[i] < 47 || emailByte[i] > 122){
			continue
		}			
		if emailByte[i] == 64 || i == 13{
			break
		}
		finalUsername = finalUsername + string(emailByte[i])
	}
	return finalUsername
}
