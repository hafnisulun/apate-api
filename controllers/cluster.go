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

type ClusterController struct{}

// GET /residences/:residence_uuid/clusters
// Get clusters by residence UUID
func (r ClusterController) FindAll(c *gin.Context) {
	// Bind query
	var query models.FindResidencesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Println("Bind request query failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Get clusters
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_RESIDENCES_BASE_URL") + "/residences/" + c.Param("residence_uuid") + "/clusters",
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

	defer res.Body.Close()
	var clustersResponseBody models.ClustersResponseBody
	json.NewDecoder(res.Body).Decode(&clustersResponseBody)

	c.JSON(http.StatusOK, clustersResponseBody)
}

// GET /residences/:residence_uuid/clusters/:cluster_uuid
// Get cluster detail
func (r ClusterController) FindOne(c *gin.Context) {
	// Get cluster
	rd := httpclient.RequestDetails{
		URL:    os.Getenv("APATE_RESIDENCES_BASE_URL") + "/residences/" + c.Param("residence_uuid") + "/clusters/" + c.Param("cluster_uuid"),
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

	defer res.Body.Close()
	var clusterResponseBody models.ClusterResponseBody
	json.NewDecoder(res.Body).Decode(&clusterResponseBody)

	c.JSON(http.StatusOK, clusterResponseBody)
}
