package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        uuid.UUID
	Name      string
	CtegoryId uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Price     int
}
