package error

import "errors"



var(
	 ErrInternalServerError = errors.New("internal server error")
	 ErrSQLError = errors.New("sql database error")
	 ErrTooManyRequests = errors.New("too many request")
	 ErrUnauthorized = errors.New("unauthorized request ")
	 ErrInvalidToken = errors.New("invalid token access")
	 ErrForbidden = errors.New("forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}