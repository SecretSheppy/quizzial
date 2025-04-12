package models

import "github.com/google/uuid"

type Quiz struct {
	QuizID       uuid.UUID `gorm:"primaryKey"`
	QuizMasterID uuid.UUID
	ShortID      string `gorm:"unique"`
	Title        string
	Sections     []Section `gorm:"foreignkey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
