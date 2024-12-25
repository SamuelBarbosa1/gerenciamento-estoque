package services

import (
	"gerenciamento-estoque/models"
	"gerenciamento-estoque/utils"
)

func CreateProduct(product models.Product) models.Product {
	return utils.AddProduct(product)
}

func ListProducts() []models.Product {
	return utils.GetProducts()
}

func ModifyStock(id int, quantity int) bool {
	return utils.UpdateStock(id, quantity)
}
