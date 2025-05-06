package e

// 400 errors
const (
	// ErrInvalidRequest : when post body, query param, or path param is invalid 
	ErrInvalidRequest = 400001 + iota

	// ErrValidateRequest : error when validating the request
	ErrValidateRequest

	// ErrDecodeRequestBody : error when decode the request body
	ErrDecodeRequestBody
)

// 404 errors
const (
	// ErrResourceNotFound : When no records corresponding to the request is found in the DB
	ErrResourceNotFound = 404001 
)

// 500 errors
const (
	// ErrInternalServer : unexpected error 
	ErrInternalServer = 500001
)