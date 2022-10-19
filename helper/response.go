package helper

import "go-assignment-2/param"

func JsonResponse(status int, errorMessage *string, payload interface{}) param.Response {
	var response param.Response
	if errorMessage != nil {
		response.Status = status
		response.ErrorMessage = errorMessage
		response.Payload = nil
	} else {
		response.Status = status
		response.ErrorMessage = nil
		response.Payload = payload
	}

	return response
}
