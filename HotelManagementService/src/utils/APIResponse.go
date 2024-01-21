package utils

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []Error     `json:"errors,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

func NewErrorResponse(message string, errors ...Error) APIResponse {
	return APIResponse{
		Status:  false,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
}
