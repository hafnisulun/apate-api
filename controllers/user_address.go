package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/middlewares"
	"github.com/hafnisulun/apate-api/models"
	httpclient "github.com/hafnisulun/apate-api/utils"
)

type UserAddressController struct{}

// GET /users/:user_uuid/addresses
// Find all addresses of authenticated user
func (r UserAddressController) FindAll(c *gin.Context) {
	// Get authenticated user
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Get user details
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String() + "/addresses",
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
	var userAddressesResponseBody models.UserAddressesResponseBody
	json.NewDecoder(res.Body).Decode(&userAddressesResponseBody)

	// Send response
	c.JSON(http.StatusOK, userAddressesResponseBody)
}

// GET /users/me/addresses/:user_address_uuid
// Find an address of authenticated user
func (r UserAddressController) FindOne(c *gin.Context) {
	// Get authenticated user
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Get user details
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String() + "/addresses/" + c.Param("user_address_uuid"),
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

	// Find address not OK
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	defer res.Body.Close()
	var userAddressResponseBody models.UserAddressResponseBody
	json.NewDecoder(res.Body).Decode(&userAddressResponseBody)

	// Send response
	c.JSON(http.StatusOK, userAddressResponseBody)
}

// POST /users/me/addresses
// Create an address of authenticated user
func (r UserAddressController) Create(c *gin.Context) {
	// Get authenticated user
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Bind input
	var input models.UserAddress
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind request body failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Create user address
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(input)
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String() + "/addresses",
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

	// Create failed
	if res.StatusCode != http.StatusCreated {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	defer res.Body.Close()
	var userAddressResponseBody models.UserAddressResponseBody
	json.NewDecoder(res.Body).Decode(&userAddressResponseBody)

	// Send response
	c.JSON(http.StatusCreated, userAddressResponseBody)
}

// PUT /users/me/addresses/:user_address_uuid
// Update an address of authenticated user
func (r UserAddressController) Update(c *gin.Context) {
	// Get authenticated user
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Bind input
	var input models.UserAddress
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind request body failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Update user address
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(input)
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String() + "/addresses/" + c.Param("user_address_uuid"),
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
	var userAddressResponseBody models.UserAddressResponseBody
	json.NewDecoder(res.Body).Decode(&userAddressResponseBody)

	// Send response
	c.JSON(http.StatusOK, userAddressResponseBody)
}

// DELETE /users/me/addresses/:user_address_uuid
// Delete an address of authenticated user
func (r UserAddressController) Delete(c *gin.Context) {
	// Get authenticated user
	authUser := c.Request.Context().Value(middlewares.ContextKeyAuthUser).(*models.AuthUser)

	// Delete user address
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_USERS_BASE_URL") + "/users/" + authUser.UUID.String() + "/addresses/" + c.Param("user_address_uuid"),
		Method: "DELETE",
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

	// Delete failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	defer res.Body.Close()
	var userAddressResponseBody models.UserAddressResponseBody
	json.NewDecoder(res.Body).Decode(&userAddressResponseBody)

	// Send response
	c.JSON(http.StatusOK, userAddressResponseBody)
}
