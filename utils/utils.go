package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// singletone validator
var Validate = validator.New()

func ParseJSONFromHttpRequest(request *http.Request, payload any) error {
	if request.Body == nil {
		return fmt.Errorf("request body is nil")
	}

	return json.NewDecoder(request.Body).Decode(payload)
}

func WriteJSONResponse(writer http.ResponseWriter, requestStatus int, responseData any) error {
	writer.Header().Add("Content-Type", "application-json")
	writer.WriteHeader(requestStatus)

	return json.NewEncoder(writer).Encode(responseData)
}

func WriteError(writer http.ResponseWriter, requestStatus int, err error) {
	WriteJSONResponse(writer, requestStatus, map[string]string{"error": err.Error()})
}
