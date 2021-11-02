package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tuanna7593/gosample/app/usecase/payload"
)

type BaseHandler struct{}

func (b *BaseHandler) SetError(w http.ResponseWriter, err error) {
	// set default writer
	b.setDefaultWriter(w)
	if err == nil {
		return
	}

	switch e := err.(type) {
	case payload.Error:
		writeErrorStatusCode(w, e.Type)
		writeErrorMessage(w, e.Error())
	case payload.Errors:
		writeErrorStatusCode(w, e[0].Type)
		writeErrorMessage(w, e.Error())
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *BaseHandler) WriteResponse(w http.ResponseWriter, statusCode int, value interface{}) {
	// set default writer
	b.setDefaultWriter(w)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		// force return 500 error
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (b *BaseHandler) setDefaultWriter(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func writeErrorStatusCode(w http.ResponseWriter, eType payload.ErrorType) {
	switch eType {
	case payload.ErrorTypeInvalidArgument, payload.ErrorTypeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case payload.ErrorTypeNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type responseErr struct {
	Error string `json:"error"`
}

func writeErrorMessage(w http.ResponseWriter, message string) {
	respErr := responseErr{
		Error: message,
	}

	if err := json.NewEncoder(w).Encode(respErr); err != nil {
		// force return 500 error
		w.WriteHeader(http.StatusInternalServerError)
	}
}
