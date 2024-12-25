package routes

import (
	"gerenciamento-estoque/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.POST("/products", handlers.CreateProductHandler)

	e.GET("/products", handlers.ListProductsHandler)

	e.PUT("/products/:id/stock/:quantity", handlers.UpdateStockHandler)
}
