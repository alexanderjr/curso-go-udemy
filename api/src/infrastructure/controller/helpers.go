package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var ErrInternalServer = errors.New("Internal Server Error")

func toJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")

	if statusCode == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func toError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("content-type", "application/json")
	toJson(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
