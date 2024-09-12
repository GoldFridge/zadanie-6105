package models

import (
	"github.com/google/uuid"
	"time"
)

// TenderStatus представляет статус тендера
// @Description TenderStatus содержит возможные статусы тендера
type TenderStatus string

const (
	TenderCreated   TenderStatus = "CREATED"
	TenderPublished TenderStatus = "PUBLISHED"
	TenderClosed    TenderStatus = "CLOSED"
)

// Tender представляет тендер
// @Description Tender содержит информацию о тендере
// @Accept json
// @Produce json
// @Success 200 {object} Tender
type Tender struct {
	ID              uuid.UUID    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name            string       `json:"name" gorm:"type:varchar(255)"`
	Description     string       `json:"description" gorm:"type:varchar(255)"`
	OrganizationID  uuid.UUID    `json:"organization_id" gorm:"type:uuid;not null"`
	CreatorUsername string       `json:"creator_username" gorm:"type:varchar(255)"`
	CreatedBy       uuid.UUID    `json:"created_by" gorm:"type:uuid;not null"` // ID пользователя
	Status          TenderStatus `json:"status" gorm:"type:varchar(20);not null;default:'CREATED'"`
	Version         int          `json:"version" gorm:"default:1"`
	CreatedAt       time.Time    `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt       time.Time    `json:"updated_at" gorm:"default:current_timestamp"`
}

func (Tender) TableName() string {
	return "tender"
}
