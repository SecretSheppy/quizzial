package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Quiz struct {
	QuizID       uuid.UUID `gorm:"primaryKey"`
	QuizMasterID uuid.UUID `gorm:"not null"`
	ShortID      string    `gorm:"not null;unique"`
	Title        string    `gorm:"not null"`
	Sections     []Section `gorm:"foreignkey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewQuiz(qmID uuid.UUID, shortID, title string) (*Quiz, error) {
	if qmID == uuid.Nil {
		return nil, fmt.Errorf("qmID cannot be nil")
	}

	if shortID == "" {
		return nil, fmt.Errorf("shortID cannot be nil")
	}

	if title == "" {
		return nil, fmt.Errorf("title cannot be nil")
	}

	q := &Quiz{
		QuizID:       uuid.New(),
		QuizMasterID: qmID,
		ShortID:      shortID,
		Title:        title,
	}

	return q, nil
}
