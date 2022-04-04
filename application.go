package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/hafnisulun/apate-api/routes"
)

func init() {
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		routes.Auth(v1)
		routes.Users(v1)
		routes.Residences(v1)
		routes.Merchants(v1)
		routes.Products(v1)
	}
	router.Run()
}
