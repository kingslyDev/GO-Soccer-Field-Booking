package error

import "errors"

var (
	ErrUserNotFound = errors.New("User Not Found")
	ErrPasswordIncorect = errors.New("Wrong Password")
	ErrUsernameExist = errors.New("Username Already Exist")
	ErrPasswordDoesNotMatch = errors.New("Password Does Not Match")

)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
	
}