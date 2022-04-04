package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/controllers"
	"github.com/hafnisulun/apate-api/middlewares"
)

func Products(router *gin.RouterGroup) {
	productsGroup := router.Group("/products").Use(middlewares.RequireJWTAuth())
	{
		product := new(controllers.ProductController)
		productsGroup.GET("", product.FindAll)
		productsGroup.GET("/:product_uuid", product.FindOne)
	}
}
