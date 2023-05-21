package response

type apiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResponse(status, message string, data interface{}) apiResponse {
	return apiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
