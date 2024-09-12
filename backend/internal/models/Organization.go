package models

import (
	"github.com/google/uuid"
	"time"
)

// OrganizationType представляет тип организации
// @Description OrganizationType содержит возможные типы организаций

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

// Organization представляет организацию
// @Description Organization содержит информацию об организации
// @Accept json
// @Produce json
// @Success 200 {object} Organization
type Organization struct {
	ID          uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string           `json:"name" gorm:"type:varchar(100);not null"`
	Description string           `json:"description" gorm:"type:text"`
	Type        OrganizationType `json:"type" gorm:"type:organization_type"`
	CreatedAt   time.Time        `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time        `json:"updated_at" gorm:"default:current_timestamp"`
}

// OrganizationResponsible представляет ответственное лицо в организации
// @Description OrganizationResponsible содержит информацию об ответственном лице
// @Accept json
// @Produce json
// @Success 200 {object} OrganizationResponsible
type OrganizationResponsible struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrganizationID uuid.UUID `json:"organization_id" gorm:"type:uuid;not null"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
}
