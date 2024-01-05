package response

import "errors"

// General error
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is invalid")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 character")
	ErrAuthIsNotExists       = errors.New("auth is not exists")
	ErrEmailAlreadyUsed       = errors.New("email already used")
)
