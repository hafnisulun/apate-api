package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Base struct {
	ID        uint       `json:"-"`
	UUID      uuid.UUID  `json:"uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
