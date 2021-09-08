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
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	HttpPortAddress, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		HttpPortAddress = 5000
	}

	ConnectionStringDatabase = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
