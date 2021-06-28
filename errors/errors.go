package errors

import "fmt"

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
type Error struct {
	status  int
	code    int
	field   string
	message string
}

func (e Error) Error() string {
	return e.message
}

func ErrOf(err error) Error {
	switch err.(type) {
	case Error:
		return err.(Error)
	default:
		return Error{status: 500, message: err.Error()}
	}
}

func ErrResp(err error) (int, Error) {
	e := ErrOf(err)
	return e.status, e
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func InvalidInputError() error {
	return Error{400, 1234, "", "Invalid input provided"}
}

func NotFoundError(id string) error {
	return Error{404, 1236, "user", fmt.Sprintf("Entity %s not found in the system", id)}
}
