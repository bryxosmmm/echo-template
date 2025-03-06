package main

import (
	_ "echo-template/docs"
	"echo-template/internal/application"
	"echo-template/internal/infrastructure"
)

func main() {
	config := infrastructure.LoadConfig()
	app := application.NewApplication(config)
	if err := app.RunServer(); err != nil {
		return
	}
}
