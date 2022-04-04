package models

import "github.com/gofrs/uuid"

type gender string

const (
	Male   gender = "male"
	Female gender = "female"
)

type User struct {
	Base
	Email     string        `json:"email"`
	Password  string        `json:"-"`
	Name      string        `json:"name"`
	Phone     *string       `json:"phone"`
	Gender    *string       `json:"gender"`
	Addresses []UserAddress `json:"-"`
}

type CreateUserInput struct {
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6,max=32"`
	Name     string  `json:"name" binding:"required,min=3,max=32"`
	Phone    *string `json:"phone"`
	Gender   *string `json:"gender"`
}

type UpdateUserInput struct {
	Name   string  `json:"name" binding:"required"`
	Phone  *string `json:"phone"`
	Gender *string `json:"gender"`
}

type AuthUser struct {
	UUID uuid.UUID `json:"uuid"`
}

type UserResponseBody struct {
	Data User `json:"data"`
}

type AuthUserResponseBody struct {
	Data AuthUser `json:"data"`
}
