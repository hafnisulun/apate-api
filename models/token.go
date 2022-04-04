package models

import "github.com/gofrs/uuid"

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   uuid.UUID
	RefreshUuid  uuid.UUID
	AtExpires    int64
	RtExpires    int64
}

type Token struct {
	UUID      uuid.UUID `json:"-"`
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	ExpiresIn float64   `json:"expires_in"`
}

type GenerateTokenRequestBody struct {
	GrantType    string `json:"grant_type" binding:"required,oneof=password refresh_token"`
	Email        string `json:"email" binding:"omitempty,email"`
	Password     string `json:"password" binding:"omitempty,min=6,max=32"`
	RefreshToken string `json:"refresh_token" binding:"required_if=GrantType refresh_token"`
}

type GenerateTokenResponseBody struct {
	AccessToken  Token `json:"access_token"`
	RefreshToken Token `json:"refresh_token"`
}
