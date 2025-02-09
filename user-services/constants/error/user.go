package error

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrPasswordIncorect = errors.New("wrong password")
	ErrUsernameExist = errors.New("username already exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")

)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,

}