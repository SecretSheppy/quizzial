package models

import (
	"errors"
	"github.com/google/uuid"
)

type QuizMaster struct {
	QuizMasterID uuid.UUID `gorm:"primaryKey"`
	Name         string    `gorm:"not null;unique"`
	Password     string    `gorm:"not null"`
	Quizzes      []Quiz    `gorm:"foreignkey:QuizMasterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewQuizMaster(name, password string) (*QuizMaster, error) {
	if name == "" {
		return nil, errors.New("invalid quiz master name")
	}

	if password == "" {
		return nil, errors.New("invalid quiz master password")
	}

	qm := &QuizMaster{
		QuizMasterID: uuid.New(),
		Name:         name,
		Password:     password,
	}

	return qm, nil
}
