package responses

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"successfully"`
	Data    interface{} `json:"data,omitempty"`
}

func Ok(code int, message string, payload interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    payload,
	}
}
