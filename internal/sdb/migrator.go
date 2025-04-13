package sdb

import (
	"github.com/SecretSheppy/quizzial/internal/models"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.SectionResource{},
		&models.QuestionResource{},
		&models.QuizMaster{},
		&models.Quiz{},
		&models.Section{})
}
