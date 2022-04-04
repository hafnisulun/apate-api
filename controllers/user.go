package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/middlewares"
	"github.com/hafnisulun/apate-api/models"
	httpclient "github.com/hafnisulun/apate-api/utils"
)

type UserController struct{}

// POST /users
// Register user
func (r UserController) Register(c *gin.Context) {
	time.Sleep(3 * time.Second)
	// Bind input
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind request body failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Create user
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(input)
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users",
		Method: "POST",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_USERS_TOKEN"),
			},
		},
		Body: payloadBuf,
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Users failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// User invalid
	if res.StatusCode != http.StatusCreated {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	defer res.Body.Close()
	var userResponseBody models.UserResponseBody
	json.NewDecoder(res.Body).Decode(&userResponseBody)

	c.JSON(http.StatusCreated, userResponseBody)
}

// GET /users/me
// Get my details
func (r UserController) FindMe(c *gin.Context) {
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Get user details
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String(),
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_USERS_TOKEN"),
			},
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Users failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	var userResponseBody models.UserResponseBody
	json.NewDecoder(res.Body).Decode(&userResponseBody)

	c.JSON(http.StatusOK, userResponseBody)
}

// PUT /users/me
// Update my details
func (r UserController) UpdateMe(c *gin.Context) {
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Bind input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind request body failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Update user details
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(input)
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String(),
		Method: "PUT",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_USERS_TOKEN"),
			},
		},
		Body: payloadBuf,
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Users failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Update failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	defer res.Body.Close()
	var userResponseBody models.UserResponseBody
	json.NewDecoder(res.Body).Decode(&userResponseBody)

	c.JSON(http.StatusOK, userResponseBody)
}
