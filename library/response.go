package library



type AppResponse struct {
	Message string `json:"mesage"`
	Data any `json:"data"`
}



func NewResponse(msg string, data any) AppResponse  {
	return AppResponse{
		Message: msg,
		Data: data,
	}
}
