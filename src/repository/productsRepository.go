package repository

import (
	"storeManager/src/config"
	"storeManager/src/models"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func GetProductById(id int) (models.Product, error) {
	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}
