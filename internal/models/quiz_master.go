package models

import "github.com/google/uuid"

type QuizMaster struct {
	QuizMasterID uuid.UUID `gorm:"primaryKey"`
	Name         string    `gorm:"unique"`
	Password     string
	Quizzes      []Quiz `gorm:"foreignkey:QuizMasterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
