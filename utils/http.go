package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload"`
}

func Response(res http.ResponseWriter, statusCode int, data *interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := response{statusCode, http.StatusText(statusCode), data}
	byteOfData, _ := json.Marshal(response)
	res.Write(byteOfData)
}

