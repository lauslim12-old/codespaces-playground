package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Status     string      `json:"status"`
	StatusText string      `json:"statusText"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewResponse(fail bool, code int, message string, data interface{}) *Response {
	if fail {
		return &Response{
			Status:     "fail",
			StatusText: http.StatusText(code),
			Code:       code,
			Message:    message,
			Data:       data,
		}
	}

	return &Response{
		Status:     "success",
		StatusText: http.StatusText(code),
		Code:       code,
		Message:    message,
		Data:       data,
	}
}

func sendResponse(w http.ResponseWriter, response *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Codespaces!"))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var object interface{}
		err := json.NewDecoder(r.Body).Decode(&object)
		if err != nil {
			sendResponse(w, NewResponse(true, http.StatusBadRequest, err.Error(), make([]interface{}, 0)))
			return
		}

		sendResponse(w, NewResponse(false, http.StatusOK, "Successfully processed POST request.", object))
	})

	log.Fatal(http.ListenAndServe(":5000", r))
}
