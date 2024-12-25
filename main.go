package main

import (
	"gerenciamento-estoque/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	routes.RegisterRoutes(e)

	// Iniciar com "go run main.go"
	e.Logger.Fatal(e.Start(":8080"))
}
