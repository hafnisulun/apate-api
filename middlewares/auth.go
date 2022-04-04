package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/models"
	httpclient "github.com/hafnisulun/apate-api/utils"
)

type contextKey string

const ContextKeyAuthUser = contextKey("auth_user")

func RequireJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user UUID
		rd := httpclient.RequestDetails{
			URL:    os.Getenv("APATE_AUTH_BASE_URL") + "/auth",
			Method: "GET",
			Headers: []httpclient.Header{
				{
					Key:   "X-Api-Token",
					Value: os.Getenv("APATE_USERS_TOKEN"),
				},
				{
					Key:   "Authorization",
					Value: "Bearer " + httpclient.ExtractAuthHeader(c.Request, "Bearer"),
				},
			},
		}
		res, err := httpclient.Send(rd)
		if err != nil {
			log.Println("[Error] Request to Apate Auth failed, err:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// User invalid
		if res.StatusCode != http.StatusOK {
			c.AbortWithStatus(res.StatusCode)
			return
		}

		defer res.Body.Close()
		var authUserResponseBody models.AuthUserResponseBody
		json.NewDecoder(res.Body).Decode(&authUserResponseBody)

		// Set authenticated user to context
		ctx := context.WithValue(c, ContextKeyAuthUser, &authUserResponseBody.Data)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
