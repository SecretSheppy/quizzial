package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Section struct {
	SectionID   uuid.UUID `gorm:"primaryKey"`
	QuizID      uuid.UUID `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string
	Cover       SectionResource `gorm:"foreignkey:SectionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Questions   []Question      `gorm:"foreignkey:SectionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// NewSection does not require a Resource as quiz masters can just leave the title pages as plain text.
func NewSection(qID uuid.UUID, title, description string) (*Section, error) {
	if qID == uuid.Nil {
		return nil, fmt.Errorf("quiz id cannot be nil")
	}

	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	s := &Section{
		SectionID:   uuid.New(),
		QuizID:      qID,
		Title:       title,
		Description: description,
	}

	return s, nil
}
