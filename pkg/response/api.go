package response

import (
	"net/http"
)

func JSON(w http.ResponseWriter, status int, msg string) {
	response := &BaseResponse{
		Message: msg,
		Status:  status,
	}
	stringResponse, _ := response.MarshalJSON()
	WriteBody(w, status, stringResponse)
}

func JSONError(w http.ResponseWriter, status int, error string) {
	response := &BaseResponse{
		Error:  error,
		Status: status,
	}
	stringResponse, _ := response.MarshalJSON()
	WriteBody(w, status, stringResponse)
}

func WriteBody(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(body)

	if err != nil {
		panic(err)
	}
}
