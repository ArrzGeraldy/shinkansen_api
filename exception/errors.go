package exception

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrServer         = errors.New("internal server error")
	ErrParameter      = errors.New("invalid value parameter")
	ErrUsernameExists = errors.New("username already exists")
	ErrLogin          = errors.New("invalid username or password")
	ErrInvalidInput   = errors.New("input field min 3 and max 100")
)