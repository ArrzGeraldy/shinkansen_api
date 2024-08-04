package helper

import (
	"encoding/json"
	"net/http"
)

func ToResponseJson(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)

	PanicIfError(err)
}