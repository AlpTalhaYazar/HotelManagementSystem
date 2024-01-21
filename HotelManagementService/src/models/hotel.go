package models

import (
	"github.com/google/uuid"
)

type Hotel struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Stars       int       `json:"stars"`
}
