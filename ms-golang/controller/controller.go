package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/caiquefds/ms-users/config"
	"github.com/caiquefds/ms-users/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(ctx *gin.Context, request any, err error) {

	var response []*model.ErrorResponse

	var jsonError *json.SyntaxError

	var unmarshalError *json.UnmarshalTypeError

	if errors.As(err, &unmarshalError) {

		unmarshalError = err.(*json.UnmarshalTypeError)
		message := config.FindErrorMessage("400.002")
		fieldError := unmarshalError.Field
		message = fmt.Sprintf(message, fieldError)

		errorReponse := &model.ErrorResponse{
			Code:    "400.002",
			Message: message,
		}

		response = append(response, errorReponse)
		ctx.JSON(http.StatusBadRequest, response)
		return

	}

	if errors.As(err, &jsonError) {
		errorReponse := &model.ErrorResponse{
			Code:    "400.000",
			Message: config.FindErrorMessage("400.000"),
		}

		response = append(response, errorReponse)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldName := fieldError.Field()

			fieldValidation := fieldError.ActualTag()
			code, message := getMessage(request, fieldName, fieldValidation)

			errorReponse := &model.ErrorResponse{
				Code:    code,
				Message: message,
			}

			response = append(response, errorReponse)
		}

		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusBadRequest, response)
		ctx.Abort()

	}

}

func getMessage(stc any, fieldName string, validation string) (messageCode string, message string) {
	requestType := reflect.TypeOf(stc)
	field, _ := requestType.FieldByName(fieldName)

	fieldJson := field.Tag.Get("json")
	messageCode = field.Tag.Get("message")
	binding := strings.Split(field.Tag.Get("binding"), ",")

	messagesCodes := strings.Split(messageCode, ",")

	messageCode = getMessageCode(messagesCodes, messageCode, validation)

	message = setMessageParams(binding, validation, message, messageCode, fieldJson)

	return messageCode, message
}

func getMessageCode(messagesCodes []string, messageCode string, validation string) string {
	for _, m := range messagesCodes {
		messageCode = strings.Split(m, "=")[0]

		test := strings.Split(m, "=")[1] == validation

		if test {
			messageCode = strings.Split(m, "=")[0]
			break
		}

	}
	return messageCode
}

func setMessageParams(binding []string, validation string, message string, messageCode string, fieldJson string) string {
	for _, b := range binding {
		valueBinding := strings.Split(b, "=")[0]
		valueBindingParams := strings.Contains(b, "=")

		if valueBinding == validation {
			if !valueBindingParams {
				message = config.FindErrorMessage(messageCode)
				message = fmt.Sprintf(message, fieldJson)
				break
			}

			if valueBindingParams {
				valueBinding = strings.Split(b, "=")[1]
				message = config.FindErrorMessage(messageCode)
				message = strings.Split(fmt.Sprintf(message, fieldJson, valueBinding), "%!")[0]
				break
			}
		}

	}
	return message
}

func createErrorMessage(err error) *model.ErrorResponse {
	messageCode := err.Error()
	message := config.FindErrorMessage(messageCode)

	errorReponse := &model.ErrorResponse{
		Code:    messageCode,
		Message: message,
	}
	return errorReponse
}
