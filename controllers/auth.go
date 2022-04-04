package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/models"
	httpclient "github.com/hafnisulun/apate-api/utils"
)

type AuthController struct{}

// POST /token
// Generate token
func (r AuthController) GenerateToken(c *gin.Context) {
	// Bind input
	var input models.GenerateTokenRequestBody
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind request body failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Generate token to Apate Auth
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(input)

	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_AUTH_BASE_URL") + "/auth",
		Method: "POST",
		Headers: []httpclient.Header{
			{
				Key:   "Token",
				Value: os.Getenv("APATE_AUTH_TOKEN"),
			},
		},
		Body: payloadBuf,
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Auth failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	var responseBody models.GenerateTokenResponseBody
	json.NewDecoder(res.Body).Decode(&responseBody)

	// Generate token failed
	if res.StatusCode != http.StatusCreated {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	c.JSON(res.StatusCode, responseBody)
}
