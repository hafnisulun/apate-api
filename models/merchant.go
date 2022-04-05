package models

import "github.com/gofrs/uuid"

type Merchant struct {
	Base
	Name        string    `json:"name"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	Phone       string    `json:"phone"`
	Image       string    `json:"image"`
	Address     string    `json:"address"`
	ClusterUUID uuid.UUID `json:"cluster_uuid"`
}

type MerchantsResponseBody struct {
	Data []Merchant `json:"data"`
	Meta Meta       `json:"meta"`
}

type MerchantResponseBody struct {
	Data Merchant `json:"data"`
}

type FindMerchantsQuery struct {
	ResidenceUUID string `form:"residence_uuid"`
	Page          int    `form:"page"`
	PerPage       int    `form:"per_page"`
}
