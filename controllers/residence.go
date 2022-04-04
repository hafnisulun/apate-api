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

type ResidenceController struct{}

// GET /residences
// Get residences
func (r ResidenceController) FindAll(c *gin.Context) {
	// Bind query
	var query models.FindResidencesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Println("Bind request query failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Get residences
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_RESIDENCES_BASE_URL") + "/residences",
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_RESIDENCES_TOKEN"),
			},
		},
		Query: map[string]string{
			"page":     strconv.Itoa(query.Page),
			"per_page": strconv.Itoa(query.PerPage),
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Residences failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get residences failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	// Get residences succeed
	defer res.Body.Close()
	var residencesResponseBody models.ResidencesResponseBody
	json.NewDecoder(res.Body).Decode(&residencesResponseBody)

	// Send response
	c.JSON(http.StatusOK, residencesResponseBody)
}

// GET /residences/:residence_uuid
// Get residence detail
func (r ResidenceController) FindOne(c *gin.Context) {
	// Get residence
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_RESIDENCES_BASE_URL") + "/residences/" + c.Param("residence_uuid"),
		Method: "GET",
		Headers: []httpclient.Header{
			{
				Key:   "X-Api-Token",
				Value: os.Getenv("APATE_RESIDENCES_TOKEN"),
			},
		},
	}
	res, err := httpclient.Send(rd)
	if err != nil {
		log.Println("[Error] Request to Apate Residences failed, err:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Get residence failed
	if res.StatusCode != http.StatusOK {
		c.AbortWithStatus(res.StatusCode)
		return
	}

	// Get residence succeed
	defer res.Body.Close()
	var residenceResponseBody models.ResidenceResponseBody
	json.NewDecoder(res.Body).Decode(&residenceResponseBody)

	// Send response
	c.JSON(http.StatusOK, residenceResponseBody)
}
