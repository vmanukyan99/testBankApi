package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(handlingResult interface{}, w http.ResponseWriter) {
	response, marshalError := json.Marshal(handlingResult)
	if marshalError != nil {
		log.Println("response marshaling error")
	}

	_, responseError := w.Write(response)
	if responseError != nil {
		log.Println("response sending error")
	}
}
