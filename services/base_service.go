package services

import (
	"edot-test/models"
	"edot-test/utils"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	goValidator "github.com/go-playground/validator/v10"
	"strings"
)

// ConvertRequest function for conversion request
func ConvertRequest(body []byte, request interface{}, identifier int64) models.Response {

	err := json.Unmarshal(body, &request)
	if err != nil {
		logs.Error("[%d] Error unmarshal from http request body to expected object request : %v", identifier, err)
		logs.Error("[%d] Request: %s", identifier, string(body))

		return models.ResponseError("Invalid Request", utils.ErrReqInvalid)
	}

	// validate request
	result := ValidateRequest(request)
	if result.Code != utils.ErrNone {
		logs.Warn("Object request not pass validation, %s", result.ErrorMessage)
		return result
	}

	return models.ResponseSuccess(nil)
}

// ValidateRequest validate a struct.
func ValidateRequest(request interface{}) models.Response {
	v := goValidator.New()
	v.RegisterValidation("required_if_field_equal", utils.RequiredIfFieldEqual)
	err := v.Struct(request)

	var errMsg []string
	if err != nil {
		for _, e := range err.(goValidator.ValidationErrors) {
			errMsg = append(errMsg, e.Field())
		}
	}

	if len(errMsg) > 0 {
		return models.ResponseError(strings.Join(errMsg, ", ")+" cannot be empty", utils.ErrReqInvalid)
	}

	return models.ResponseSuccess(nil)
}
