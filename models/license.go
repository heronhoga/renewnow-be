package models

import (
	"time"

	"github.com/google/uuid"
)

type License struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LicenseType string    `json:"license_type"`
	Expire      time.Time `json:"expire"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
}