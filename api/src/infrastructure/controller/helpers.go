package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

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
