package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/controllers"
	"github.com/hafnisulun/apate-api/middlewares"
)

func Merchants(router *gin.RouterGroup) {
	merchantsGroup := router.Group("/merchants").Use(middlewares.RequireJWTAuth())
	{
		merchant := new(controllers.MerchantController)
		merchantsGroup.GET("", merchant.FindAll)
		merchantsGroup.GET("/:merchant_uuid", merchant.FindOne)
	}
}
