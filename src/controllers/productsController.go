package controllers

import (
	"net/http"
	"storeManager/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, erro := services.GetProducts()
	if erro != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": erro.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id, erro := strconv.ParseInt(c.Param("id"), 10, 32)
	if erro != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "O parametro precisa ser um número",
		})
		return
	}
	product, erro := services.GetProductById(int(id))
	if erro != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": erro.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}
