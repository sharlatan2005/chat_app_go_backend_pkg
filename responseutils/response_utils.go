package responseutils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Success: status < 400,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   err,
	})
}

func BadRequest(w http.ResponseWriter, err string) {
	Error(w, http.StatusBadRequest, err)
}

func Unauthorized(w http.ResponseWriter, err string) {
	Error(w, http.StatusUnauthorized, err)
}

func InternalServerError(w http.ResponseWriter, err string) {
	Error(w, http.StatusInternalServerError, err)
}

func NotFound(w http.ResponseWriter, err string) {
	Error(w, http.StatusNotFound, err)
}

func StatusMethodNotAllowed(w http.ResponseWriter, err string) {
	Error(w, http.StatusMethodNotAllowed, err)
}

func StatusBadRequest(w http.ResponseWriter, err string) {
	Error(w, http.StatusBadRequest, err)
}
