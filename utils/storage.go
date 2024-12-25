package utils

import "gerenciamento-estoque/models"

var Products = []models.Product{}

var nextID = 1

func AddProduct(product models.Product) models.Product {
	product.ID = nextID
	nextID++
	Products = append(Products, product)
	return product
}

func GetProducts() []models.Product {
	return Products
}

func UpdateStock(id int, quantity int) bool {
	for i, p := range Products {
		if p.ID == id {
			Products[i].Stock += quantity
			return true
		}
	}
	return false
}
