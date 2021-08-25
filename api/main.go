package main

import (
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Rodando a API")
	r := router.Gerar()

	http.ListenAndServe(":5000", r)
}
