package services

import (
	"storeManager/src/models"
	"storeManager/src/repository"
)

func GetProducts() ([]models.Product, error) {
	products, erro := repository.GetProducts()
	if erro != nil {
		return nil, erro
	}
	return products, nil
}

func GetProductById(id int) (models.Product, error) {
	product, erro := repository.GetProductById(id)
	if erro != nil {
		return models.Product{}, erro
	}
	return product, nil
}
