package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BuyerID string    `gorm:"type:uuid;not null" json:"buyer_id"`
	Buyer   User      `gorm:"foreignkey:BuyerID;references:ID" json:"buyer"`
	EventID string    `gorm:"type:uuid;not null" json:"event_id"`
	Event   Event     `gorm:"foreignkey:EventID;references:ID" json:"event"`
	Amount  int       `gorm:"not null" json:"amount"`
	Timestamp
}

func (e *Transaction) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
