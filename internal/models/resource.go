package models

import "github.com/google/uuid"

// Resource is not a table, it is just used as an include for other Resource models
type Resource struct {
	ResourceID uuid.UUID `gorm:"primaryKey"`
	Path       string
	Alt        string
}
