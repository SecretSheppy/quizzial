package models

import (
	"errors"
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/google/uuid"
)

type Question struct {
	QuestionID       uuid.UUID `gorm:"primaryKey"`
	SectionID        uuid.UUID
	QuestionableType string
	QuestionableID   uuid.UUID
}

func NewQuestion(sID uuid.UUID, question qplugins.QPluginModel) (*Question, error) {
	if sID == uuid.Nil {
		return nil, errors.New("sID cannot be nil")
	}

	if question == nil {
		return nil, errors.New("question cannot be nil")
	}

	q := &Question{
		QuestionID:       uuid.New(),
		SectionID:        sID,
		QuestionableType: question.GetType(),
		QuestionableID:   question.GetID(),
	}

	return q, nil
}
