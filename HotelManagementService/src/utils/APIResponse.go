package utils

import "google.golang.org/protobuf/types/known/emptypb"

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []Error     `json:"errors"`
}

func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  true,
		Message: message,
		Data:    data,
		Errors:  make([]Error, 0),
	}
}

func NewErrorResponse(message string, errors ...Error) APIResponse {
	return APIResponse{
		Status:  false,
		Message: message,
		Data:    emptypb.Empty{},
		Errors:  errors,
	}
}
