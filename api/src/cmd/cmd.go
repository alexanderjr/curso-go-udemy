package cmd

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Execute() {
	app := &cli.App{
		Name:  "Starting Users API",
		Usage: "API who is resposible for our users",
		Action: func(c *cli.Context) error {
			config.Carregar()
			logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.Info(
				fmt.Sprintf("Rodanado a API na porta %d", config.Porta),
			)

			if err := http.ListenAndServe(
				fmt.Sprintf(":%d", config.Porta),
				router.Gerar(),
			); err != nil {
				panic(err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
