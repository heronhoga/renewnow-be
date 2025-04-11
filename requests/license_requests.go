package requests

import (
	"time"
	"github.com/google/uuid"
)

type CreateLicenseRequest struct {
	LicenseType string    `json:"license_type" validate:"required"`
	Expire      time.Time `json:"expire" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required,uuid4"`
}