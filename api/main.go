package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.Carregar()
	log.SetFormatter(&log.JSONFormatter{})
	log.Info(
		fmt.Sprintf("Rodando a API na porta %d", config.Porta),
	)

	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", config.Porta),
		router.Gerar(),
	); err != nil {
		panic(err)
	}
}
