package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/controllers"
	"github.com/hafnisulun/apate-api/middlewares"
)

func Users(router *gin.RouterGroup) {
	user := new(controllers.UserController)
	userAddress := new(controllers.UserAddressController)

	usersGroup := router.Group("/users")
	{
		usersGroup.POST("", user.Register)
	}

	authUsersGroup := router.Group("/users").Use(middlewares.RequireJWTAuth())
	{
		authUsersGroup.GET("me", user.FindMe)
		authUsersGroup.PUT("me", user.UpdateMe)
		authUsersGroup.GET("me/addresses", userAddress.FindAll)
		authUsersGroup.GET("me/addresses/:user_address_uuid", userAddress.FindOne)
		authUsersGroup.POST("me/addresses", userAddress.Create)
		authUsersGroup.PUT("me/addresses/:user_address_uuid", userAddress.Update)
		authUsersGroup.DELETE("me/addresses/:user_address_uuid", userAddress.Delete)
	}
}
