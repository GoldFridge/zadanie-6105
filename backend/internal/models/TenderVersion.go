package models

import (
	"github.com/google/uuid"
	"time"
)

type TenderVersion struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TenderID    uuid.UUID `json:"tender_id" gorm:"type:uuid;not null"`
	Version     int       `json:"version" gorm:"not null"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
}

func (TenderVersion) TableName() string {
	return "tender_version"
}
