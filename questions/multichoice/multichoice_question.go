package multichoice

import (
	"github.com/google/uuid"
)

type Option struct {
	OptionID   uuid.UUID `gorm:"primaryKey"`
	QuestionID uuid.UUID
	Text       string
}

type MultiChoiceQuestion struct {
	MultiChoiceQuestionID uuid.UUID `gorm:"primaryKey"`
	Text                  string
	Options               []Option `gorm:"foreignKey:QuestionID"`
}

func NewMultiChoiceQuestion() *MultiChoiceQuestion {
	return &MultiChoiceQuestion{}
}

func (m *MultiChoiceQuestion) GetType() string {
	return "MultiChoiceQuestion"
}

func (m *MultiChoiceQuestion) GetID() uuid.UUID {
	return m.MultiChoiceQuestionID
}
