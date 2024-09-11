package models

import (
	"github.com/google/uuid"
	"time"
)

type TenderStatus string

const (
	TenderCreated   TenderStatus = "CREATED"
	TenderPublished TenderStatus = "PUBLISHED"
	TenderClosed    TenderStatus = "CLOSED"
)

type Tender struct {
	ID             uuid.UUID    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizationID uuid.UUID    `json:"organization_id" gorm:"type:uuid;not null"`
	CreatedBy      uuid.UUID    `json:"created_by" gorm:"type:uuid;not null"` // ID пользователя
	Status         TenderStatus `json:"status" gorm:"type:varchar(20);not null;default:'CREATED'"`
	Version        int          `json:"version" gorm:"default:1"`
	CreatedAt      time.Time    `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt      time.Time    `json:"updated_at" gorm:"default:current_timestamp"`
}
