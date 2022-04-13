package helper

import "os"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}
type Response struct {
	ApiVersion string      `json:"api_version"`
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		ApiVersion: os.Getenv("API_VERSION"),
		Meta:       meta,
		Data:       data,
	}
	return jsonResponse
}
