package models

type QuestionResource struct {
	Resource
	QuestionID string `gorm:"foreignKey:QuestionID"`
}
