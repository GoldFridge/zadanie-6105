package models

import (
	"github.com/google/uuid"
	"time"
)

type BidVersion struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	BidID       uuid.UUID `json:"bid_id" gorm:"type:uuid;not null"`
	Version     int       `json:"version" gorm:"not null"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	Status      BidStatus `json:"status" gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
}

// TableName задает название таблицы для модели BidVersion
func (BidVersion) TableName() string {
	return "bid_versions"
}
