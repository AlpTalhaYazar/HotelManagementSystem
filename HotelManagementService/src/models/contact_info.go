package models

import (
	"github.com/google/uuid"
)

type ContactInfo struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	HotelID uuid.UUID `gorm:"type:uuid;not null" json:"hotel_id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
}
