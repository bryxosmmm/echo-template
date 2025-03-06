package utils

import "net/http"

type Err struct {
	Message string `json:"message"`
}

func newError(msg string) *Err {
	return &Err{
		Message: msg,
	}
}

func BadRequestError() (int, *Err) {
	return http.StatusBadRequest, newError("Bad Request")
}

func NotFoundError() (int, *Err) {
	return http.StatusNotFound, newError("Not Found")
}

func InternalServerError(msg string) (int, *Err) {
	return http.StatusInternalServerError, newError(msg)
}

func (e *Err) Error() string {
	return e.Message
}

func ConflictError() (int, *Err) {
	return http.StatusConflict, newError("Some fields can't be modified")
}

func MultipleLoginError() *Err {
	return &Err{
		Message: "login should be unique",
	}
}
