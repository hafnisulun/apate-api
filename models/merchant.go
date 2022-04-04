package models

import "github.com/gofrs/uuid"

type Merchant struct {
	Base
	Name        string    `json:"name"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	Image       string    `json:"image"`
	ClusterUUID uuid.UUID `json:"cluster_uuid"`
}

type MerchantsResponseBody struct {
	Data []Merchant `json:"data"`
	Meta Meta       `json:"meta"`
}

type MerchantResponseBody struct {
	Data Merchant `json:"data"`
}

type CreateMerchantInput struct {
	Name        string    `json:"name" binding:"required"`
	Lat         float64   `json:"lat" binding:"required"`
	Lon         float64   `json:"lon" binding:"required"`
	ClusterUUID uuid.UUID `json:"cluster_uuid" binding:"required"`
}

type FindMerchantsQuery struct {
	ResidenceUUID string `form:"residence_uuid"`
	Page          int    `form:"page"`
	PerPage       int    `form:"per_page"`
}
