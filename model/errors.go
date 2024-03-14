package forum

import "errors"

// User Error Messages
var (
	UserNameError       = errors.New("User Name error!")
	UserNameLengthError = errors.New("Username must be between 3 and 13 characters")
	UserNameCharError   = errors.New("Username must contain numbers and letters only")
	UserEmailError      = errors.New("Email does not exist!")
	UserPasswordError   = errors.New("User Password error!")
	RegPasswordError    = errors.New("Password too weak!\nmust be more than 6 characters")
	UserExistsError     = errors.New("Email Already in Use!")
	UsernameExistsError = errors.New("Username Already in Use!")
	ActiveUserError     = errors.New("User Data Failed to load")
	SessionExpired      = errors.New("Session Expired")
	SessionInvalid      = errors.New("Session Does  Not Exist")
	EmailFormatError    = errors.New("Wrong Email Format!")
	NewCategoryError    = errors.New("Error adding category!")
	ScanCategoryError   = errors.New("Category Scan Error!")
	PostError           = errors.New("Post Title and Body must contain characters")
	CategoryNotSelected = errors.New("You must Select at least one Category for your post!")
)
