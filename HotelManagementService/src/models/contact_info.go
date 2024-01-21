package models

import (
	"github.com/google/uuid"
)

type ContactInfo struct {
	Id      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:Id" json:"id"`
	HotelId uuid.UUID `gorm:"type:uuid;not null;column:HotelId" json:"hotel_id"`
	Name    string    `gorm:"column:Name" json:"name"`
	Phone   string    `gorm:"column:Phone" json:"phone"`
	Email   string    `gorm:"column:Email" json:"email"`
}

func (ContactInfo) TableName() string {
	return "ContactInfos"
}
