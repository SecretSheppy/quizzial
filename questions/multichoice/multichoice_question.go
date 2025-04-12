package multichoice

import (
	"github.com/SecretSheppy/quizzial/internal/models"
	"github.com/google/uuid"
)

type Option struct {
	OptionID   uuid.UUID `gorm:"primaryKey"`
	QuestionID uuid.UUID
	Text       string
}

type MultiChoiceQuestion struct {
	models.Question
	Options []Option `gorm:"foreignKey:QuestionID"`
}

func (m *MultiChoiceQuestion) GetType() string {
	return "MultiChoiceQuestion"
}
