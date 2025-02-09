package error

import "errors"

const(
	Success = "success"
	Error = "error"
)

var(
	 ErrInternalServerError = errors.New("Internal Server Error")
	 ErrSQLError = errors.New("Sql Database Error")
	 ErrTooManyRequests = errors.New("Too many request")
	 ErrUnauthorized = errors.New("Unauthorized Request ")
	 ErrInvalidToken = errors.New("Invalid Token Access")
	 ErrForbidden = errors.New("Forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}