package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-api/models"
	httpclient "github.com/hafnisulun/apate-api/utils"
)

type MerchantController struct{}

// GET /merchants
// Get merchants
func (r MerchantController) FindAll(c *gin.Context) {
	// Bind query
	var query models.FindMerchantsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Println("Bind request query failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Get merchants
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_MERCHANTS_BASE_URL") + "/merchants",
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_MERCHANTS_TOKEN"),
			},
		},
		Query: map[string]string{
			"residence_uuid": query.ResidenceUUID,
			"page":           strconv.Itoa(query.Page),
			"per_page":       strconv.Itoa(query.PerPage),
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Merchants failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get merchants failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	// Get merchants succeed
	defer res.Body.Close()
	var merchantsResponseBody models.MerchantsResponseBody
	json.NewDecoder(res.Body).Decode(&merchantsResponseBody)

	// Send response
	c.JSON(http.StatusOK, merchantsResponseBody)
}

// GET /merchants/:merchant_uuid
// Get merchant details
func (r MerchantController) FindOne(c *gin.Context) {
	// Get merchant
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_MERCHANTS_BASE_URL") + "/merchants/" + c.Param("merchant_uuid"),
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_MERCHANTS_TOKEN"),
			},
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Merchants failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get merchant failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	// Get merchant succeed
	defer res.Body.Close()
	var merchantResponseBody models.MerchantResponseBody
	json.NewDecoder(res.Body).Decode(&merchantResponseBody)

	// Send response
	c.JSON(http.StatusOK, merchantResponseBody)
}
