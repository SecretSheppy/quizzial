package models

type SectionResource struct {
	Resource
	SectionID string `gorm:"foreignKey:SectionID"`
}
