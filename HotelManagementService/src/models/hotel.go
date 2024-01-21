package models

import (
	"github.com/google/uuid"
)

type Hotel struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:Id" json:"id"`
	Name        string    `gorm:"column:Name" json:"name"`
	Description string    `gorm:"column:Description" json:"description"`
	City        string    `gorm:"column:City" json:"city"`
	Country     string    `gorm:"column:Country" json:"country"`
	Stars       int       `gorm:"column:Stars" json:"stars"`
}

func (Hotel) TableName() string {
	return "Hotels"
}
