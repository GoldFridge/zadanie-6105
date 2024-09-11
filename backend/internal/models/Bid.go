package models

import (
	"github.com/google/uuid"
	"time"
)

type ProposalStatus string

const (
	ProposalCreated   ProposalStatus = "CREATED"
	ProposalPublished ProposalStatus = "PUBLISHED"
	ProposalCanceled  ProposalStatus = "CANCELED"
)

type Proposal struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TenderID  uuid.UUID      `json:"tender_id" gorm:"type:uuid;not null"`
	CreatedBy uuid.UUID      `json:"created_by" gorm:"type:uuid;not null"` // ID пользователя
	Status    ProposalStatus `json:"status" gorm:"type:varchar(20);not null;default:'CREATED'"`
	Version   int            `json:"version" gorm:"default:1"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
}
