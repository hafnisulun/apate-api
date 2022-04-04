package models

import "github.com/gofrs/uuid"

type Product struct {
	Base
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	Description  *string   `json:"description"`
	Image        *string   `json:"image"`
	MerchantUUID uuid.UUID `json:"merchant_uuid"`
}

type ProductsResponseBody struct {
	Data []Product `json:"data"`
	Meta Meta      `json:"meta"`
}

type ProductResponseBody struct {
	Data Product `json:"data"`
}

type CreateProducttInput struct {
	Name         string    `json:"name" binding:"required"`
	Price        float64   `json:"price" binding:"required"`
	Description  *string   `json:"description"`
	Image        *string   `json:"image"`
	MerchantUUID uuid.UUID `json:"merchant_uuid" binding:"required"`
}

type FindProductsQuery struct {
	MerchantUUID string `form:"merchant_uuid"`
	Page         int    `form:"page"`
	PerPage      int    `form:"per_page"`
}
