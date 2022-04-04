package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/controllers"
	"github.com/hafnisulun/apate-api/middlewares"
)

func Residences(router *gin.RouterGroup) {
	residencesGroup := router.Group("/residences").Use(middlewares.RequireJWTAuth())
	{
		residence := new(controllers.ResidenceController)
		cluster := new(controllers.ClusterController)
		residencesGroup.GET("", residence.FindAll)
		residencesGroup.GET("/:residence_uuid", residence.FindOne)
		residencesGroup.GET("/:residence_uuid/clusters", cluster.FindAll)
		residencesGroup.GET("/:residence_uuid/clusters/:cluster_uuid", cluster.FindOne)
	}
}
