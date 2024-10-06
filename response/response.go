package response

import (
	"encoding/json"
	"net/http"
	"time"
)

func Ok(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

func BadRequest(w http.ResponseWriter, msg error) {
	if err := json.NewEncoder(w).Encode(newErrorHandler(msg, http.StatusBadRequest)); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
	}
}

type errorHandler struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Code      int       `json:"code"`
}

func newErrorHandler(err error, statusCode int) errorHandler {
	return errorHandler{
		Message:   err.Error(),
		Code:      statusCode,
		Timestamp: time.Now(),
	}
}
