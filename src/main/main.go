package main

import (
	"log"
	"storeManager/src/config"
	"storeManager/src/models"
	"storeManager/src/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	config.ConnectDB()
	if erro := config.DB.AutoMigrate(&models.Product{}, &models.Sale{}, &models.SaleProduct{}); erro != nil {
		log.Fatal(erro)
	}
}

func main() {
	Router := gin.Default()
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routers.RoteamentoUsuarios(Router)
	Router.Run()
}
