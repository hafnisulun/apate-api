package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/controllers"
	"github.com/hafnisulun/apate-api/middlewares"
)

func Auth(router *gin.RouterGroup) {
	authGroup := router.Group("/auth").Use(middlewares.RequireClientAuth())
	{
		auth := new(controllers.AuthController)
		authGroup.POST("", auth.GenerateToken)
	}
}
