package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {

	log.Info("Including required packages and building a basic CLI")

	app := cli.NewApp()

	app.Run(os.Args)
}
