package main

import (
	_ "echo-template/docs"
	"echo-template/internal/application"
	"echo-template/internal/infrastructure"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @BasePath  /api/v1
func main() {
	config := infrastructure.LoadConfig()
	app := application.NewApplication(config)
	if err := app.RunServer(); err != nil {
		return
	}
}
