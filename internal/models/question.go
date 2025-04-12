package models

import "github.com/google/uuid"

type Question struct {
	QuestionID       uuid.UUID `gorm:"primaryKey"`
	SectionID        uuid.UUID
	Text             string
	QuestionableType string
	QuestionableID   uuid.UUID
}
