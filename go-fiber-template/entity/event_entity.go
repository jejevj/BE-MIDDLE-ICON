package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `json:"name"`
	AuthorID    uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
	Author      User      `gorm:"foreignkey:AuthorID;references:ID;constraint:OnDelete:CASCADE;" json:"author"`
	Price       int       `json:"price"`
	Capacity    int       `json:"capacity"`
	Availabilty int       `json:"availabilty"`

	Timestamp
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
