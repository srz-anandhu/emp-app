package e

import (
	"net/http"
	"strconv"
)

type WrapError struct {
	ErrorCode int
	Msg       string
	RootCause error
}

type HTTPError struct {
	StatusCode int
	Code       int
	Message    string
}

func (e *WrapError) Error() string {
	return e.RootCause.Error()
}

// NewError : creates a new error instance, get rootcause error and return as WrapError.
func NewError(errCode int, msg string, rootCause error) *WrapError {
	err := &WrapError{
		ErrorCode: errCode,
		Msg:       msg,
		RootCause: rootCause,
	}

	return err
}

// NewApiError : creates HTTP error from NewError function to pass to response.Fail
func NewApiError(err error, msg string) *HTTPError {
	if err == nil {
		return nil
	}

	// checking err is type of WrapError
	appErr, ok := err.(*WrapError)
	if ok {
		appErr.Msg = msg
	} else {
		return nil
	}

	httpError := &HTTPError{
		StatusCode: GetHttpStatusCode(appErr.ErrorCode),
		Code:       appErr.ErrorCode,
		Message:    msg,
	}

	return httpError
}

func GetHttpStatusCode(c int) int {
	str := strconv.Itoa(c)
	// Geting first 3 digints from ErrorCode (eg : 400001 => 400)
	code := str[:3]

	r, _ := strconv.Atoi(code)
	if r < 100 || r > 600 {
		return http.StatusInternalServerError
	}
	return r
}
