package models

import "github.com/google/uuid"

type Section struct {
	SectionID   uuid.UUID `gorm:"primaryKey"`
	QuizID      uuid.UUID
	Title       string
	Description string
	Cover       Resource   `gorm:"polymorphic:Resourcable;foreignkey:SectionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Questions   []Question `gorm:"polymorphic:Questionable;foreignkey:SectionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
