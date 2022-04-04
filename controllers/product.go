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

type ProductController struct{}

// GET /products
// Get products
func (r ProductController) FindAll(c *gin.Context) {
	// Bind query
	var query models.FindProductsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Println("Bind request query failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Get products
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_PRODUCTS_BASE_URL") + "/products",
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_PRODUCTS_TOKEN"),
			},
		},
		Query: map[string]string{
			"merchant_uuid": query.MerchantUUID,
			"page":          strconv.Itoa(query.Page),
			"per_page":      strconv.Itoa(query.PerPage),
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Products failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get products failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	// Get products succeed
	defer res.Body.Close()
	var productsResponseBody models.ProductsResponseBody
	json.NewDecoder(res.Body).Decode(&productsResponseBody)

	// Send response
	c.JSON(http.StatusOK, productsResponseBody)
}

// GET /products/:product_uuid
// Get product details
func (r ProductController) FindOne(c *gin.Context) {
	// Get product
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_PRODUCTS_BASE_URL") + "/products/" + c.Param("product_uuid"),
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_PRODUCTS_TOKEN"),
			},
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Products failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get product failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	// Get product succeed
	defer res.Body.Close()
	var productResponseBody models.ProductResponseBody
	json.NewDecoder(res.Body).Decode(&productResponseBody)

	// Send response
	c.JSON(http.StatusOK, productResponseBody)
}
