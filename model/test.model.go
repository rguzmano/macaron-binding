package model

import (
	"encoding/json"
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"labs/macaron-binding/util"
	"net/http"
)

type TestRequest struct {
	StringField string `json:"stringField" binding:"Required;Email;GmailValidation"`
	NumberField int    `json:"numberField" binding:"Range(1,10)"`
}

type TestGetRequest struct {
	Field1 string `form:"field1" binding:"Required"`
}

func (cf TestRequest) Error(ctx *macaron.Context, errs binding.Errors) {
	var errorResponse []ErrorResponse
	hasError := false
	for _, err := range errs {
		hasError = true
		errorResponse = append(errorResponse,
			ErrorResponse{
				Type:    err.Classification,
				Message: util.SetMessageByClassification(err.Classification, err.FieldNames[0]),
			})
	}
	if hasError {
		ctx.Resp.WriteHeader(http.StatusBadRequest)
		bResponse, _ := json.Marshal(errorResponse)
		_, _ = ctx.Resp.Write(bResponse)
		return
	}
}
