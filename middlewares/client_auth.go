package middlewares

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	httpclient "github.com/hafnisulun/apate-api/utils"
)

func RequireClientAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authByte, _ := base64.StdEncoding.DecodeString(httpclient.ExtractAuthHeader(c.Request, "Basic"))
		authString := string(authByte)
		credentials := strings.Split(authString, ":")
		if credentials[0] != os.Getenv("CLIENT_ID") || credentials[1] != os.Getenv("CLIENT_SECRET") {
			log.Println("[Error] client ID and secret invalid")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
