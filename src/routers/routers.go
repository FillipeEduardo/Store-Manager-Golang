package routers

import (
	"storeManager/src/controllers"

	"github.com/gin-gonic/gin"
)

func RoteamentoUsuarios(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
}
