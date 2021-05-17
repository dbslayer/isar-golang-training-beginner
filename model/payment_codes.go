package model

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type PaymentCodes struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PaymentCode    string    `json:"payment_code" validate:"required"`
	Name           string    `json:"name" validate:"required"`
	Status         string    `json:"status"`
	ExpirationDate string    `json:"expiration_date" gorm:"column:expiration_date"`
	CreatedDate    time.Time `json:"created_date" gorm:"column:created_date"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (paymentCodes *PaymentCodes) BeforeCreate(tx *gorm.DB) (err error) {
	paymentCodes.ID = uuid.New()
	return
}
