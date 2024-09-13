package models

import (
	"github.com/google/uuid"
	"time"
)

type BidStatus string

// BidStatus представляет статус предложения
// @Description Статусы предложения
// @Enum CREATED PUBLISHED CANCELED
const (
	BidCreated   BidStatus = "CREATED"
	BidPublished BidStatus = "PUBLISHED"
	BidCanceled  BidStatus = "CANCELED"
)

// Bid представляет собой предложение для тендера
// @Description Предложение для тендера
// @Accept json
// @Produce json
// @Success 200 {object} Proposal
type Bid struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	TenderID    uuid.UUID `json:"tender_id" gorm:"type:uuid;not null"`
	CreatedBy   uuid.UUID `json:"created_by" gorm:"type:uuid;not null"` // ID пользователя
	Status      BidStatus `json:"status" gorm:"type:varchar(20);not null;default:'CREATED'"`
	Version     int       `json:"version" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}

func (Bid) TableName() string {
	return "bid"
}
