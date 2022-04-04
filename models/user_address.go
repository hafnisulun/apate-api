package models

import "github.com/gofrs/uuid"

type UserAddress struct {
	Base
	UserID        uint      `json:"-"`
	User          User      `json:"-"`
	Label         string    `json:"label" binding:"required"`
	ResidenceUUID uuid.UUID `json:"residence_uuid" binding:"required"`
	ClusterUUID   uuid.UUID `json:"cluster_uuid" binding:"required"`
	Details       string    `json:"details" binding:"required"`
}

type UserAddressesResponseBody struct {
	Data []UserAddress `json:"data"`
	Meta Meta          `json:"meta"`
}

type UserAddressResponseBody struct {
	Data UserAddress `json:"data"`
}

// type UserAddressInput struct {
// 	Name          string    `json:"name" binding:"required"`
// 	ResidenceUUID uuid.UUID `json:"residence_uuid" binding:"required"`
// 	ClusterUUID   uuid.UUID `json:"cluster_uuid" binding:"required"`
// 	Details       string    `json:"details" binding:"required"`
// }
