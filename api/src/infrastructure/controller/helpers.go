package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var ErrInternalServer = errors.New("Internal Server Error")

func toJson(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent {
		return
	}

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func toError(w http.ResponseWriter, statusCode int, erro error) {
	toJson(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
