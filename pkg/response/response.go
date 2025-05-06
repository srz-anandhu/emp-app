package response

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClentResponse(status int, message string, data, err interface{}) Response {
	return Response{
		StatusCode: status,
		Message: message,
		Data: data,
		Error: err,
	}
}