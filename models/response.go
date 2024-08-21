package models

import "edot-test/utils"

// Response struct
type Response struct {
	Content        interface{} `json:"content"`
	SuccessMessage string      `json:"successMessage"`
	Code           int         `json:"code"`
	ErrorMessage   string      `json:"errorMessage"`
}

// ResponseSuccess function to generate non pointer success response
func ResponseSuccess(result interface{}) Response {
	return Response{
		Code:           utils.ErrNone,
		Content:        result,
		SuccessMessage: "Success",
	}
}

// ResponseError function to generate non pointer error response
func ResponseError(err string, code int) Response {
	return Response{
		Code:         code,
		ErrorMessage: err,
	}
}
