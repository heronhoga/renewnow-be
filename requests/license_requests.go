package requests

import (
	"time"
)

type CreateLicenseRequest struct {
	LicenseType string    `json:"license_type" validate:"required"`
	Expire      time.Time `json:"expire" validate:"required"`
}