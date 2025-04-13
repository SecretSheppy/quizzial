package multichoice

import (
	"errors"
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/google/uuid"
)

type Option struct {
	OptionID              uuid.UUID `gorm:"primaryKey"`
	MultiChoiceQuestionID uuid.UUID
	Text                  string
	Correct               bool
}

type MultiChoiceQuestion struct {
	MultiChoiceQuestionID uuid.UUID `gorm:"primaryKey"`
	Text                  string
	Options               []Option `gorm:"foreignKey:MultiChoiceQuestionID"`
}

func NewMultiChoiceQuestion(text string) (*MultiChoiceQuestion, error) {
	if text == "" {
		return nil, errors.New("empty question text")
	}

	mcq := &MultiChoiceQuestion{
		MultiChoiceQuestionID: uuid.New(),
		Text:                  text,
	}

	return mcq, nil
}

func (m *MultiChoiceQuestion) New() qplugins.QPluginModel {
	return &MultiChoiceQuestion{}
}

func (m *MultiChoiceQuestion) GetType() string {
	return "MultiChoiceQuestion"
}

func (m *MultiChoiceQuestion) GetID() uuid.UUID {
	return m.MultiChoiceQuestionID
}
