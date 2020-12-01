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

type responseNoPayload struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
}

func Response(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := response{statusCode, http.StatusText(statusCode), data}
	byteOfData, _ := json.Marshal(response)
	res.Write(byteOfData)
}

func ResponseWithoutPayload(res http.ResponseWriter, statusCode int){
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := responseNoPayload{statusCode, http.StatusText(statusCode)}
	byteOfData, _ := json.Marshal(response)
	res.Write(byteOfData)
}

func ResponseCustom(res http.ResponseWriter, statusCode int,message string, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := response{statusCode, message, data}
	byteOfData, _ := json.Marshal(response)
	res.Write(byteOfData)
}
