package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionStringDatabase = ""
	HttpPortAddress          = 0
)

func LoadVars() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	HttpPortAddress, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		HttpPortAddress = 9000
	}

	ConnectionStringDatabase = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
