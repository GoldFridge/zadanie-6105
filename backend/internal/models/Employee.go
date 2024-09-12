package models

import (
	"github.com/google/uuid"
	"time"
)

// Employee представляет собой сотрудника организации
// @Description Сотрудник организации
// @Accept json
// @Produce json
// @Success 200 {object} Employee
type Employee struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string    `json:"username" gorm:"type:varchar(50);unique;not null"`
	FirstName string    `json:"first_name" gorm:"type:varchar(50)"`
	LastName  string    `json:"last_name" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
