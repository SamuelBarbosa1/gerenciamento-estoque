package handlers

import (
	"gerenciamento-estoque/models"
	"gerenciamento-estoque/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductHandler(c echo.Context) error {

	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	createdProduct := services.CreateProduct(*product)
	return c.JSON(http.StatusCreated, createdProduct)
}

func ListProductsHandler(c echo.Context) error {
	products := services.ListProducts()
	return c.JSON(http.StatusOK, products)
}

func UpdateStockHandler(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	quantity, err := strconv.Atoi(c.Param("quantity"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Quantidade inválida"})
	}

	success := services.ModifyStock(id, quantity)
	if !success {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Estoque atualizado com sucesso"})
}
