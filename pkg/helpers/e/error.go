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
	if e == nil {
		return "<nil>"
	}
	if e.RootCause != nil {
		return e.Msg + ": " + e.RootCause.Error()
	}
	return e.Msg
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

	var errorCode int
	var errorMsg string

	if wrapErr, ok := err.(*WrapError); ok {
		errorCode = wrapErr.ErrorCode
		if msg != "" {
			errorMsg = msg
		} else {
			errorMsg = wrapErr.Msg
		}
	} else {
		errorCode = ErrInternalServer 
		errorMsg = err.Error()
		if msg != "" {
			errorMsg = msg + ": " + errorMsg
		}
	}
	httpError := &HTTPError{
		StatusCode: GetHttpStatusCode(errorCode),
		Code:       errorCode,
		Message:    errorMsg,
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
