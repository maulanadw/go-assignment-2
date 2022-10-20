package helper

import "go-assignment-2/param"

func JsonResponse(status int, message *string, errorInfo *string, payload interface{}) param.Response {
	var response param.Response
	if errorInfo != nil {
		response.Status = status
		response.Message = message
		response.ErrorInfo = errorInfo
		response.Payload = nil
	} else {
		response.Status = status
		response.Message = message
		response.ErrorInfo = nil
		response.Payload = payload
	}

	return response
}
