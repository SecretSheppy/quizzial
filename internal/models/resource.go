package models

import "github.com/google/uuid"

type Resource struct {
	ResourceID      uuid.UUID `gorm:"primaryKey"`
	Path            string
	Alt             string
	ResourcableType string
	ResourcableID   uuid.UUID
}
