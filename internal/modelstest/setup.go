package modelstest

import (
	"errors"
	"github.com/SecretSheppy/quizzial/internal/models"
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.QuizMaster{},
		&models.Quiz{},
		&models.Section{},
		&models.SectionResource{},
		&models.Question{})
}

func CreateTestUser(db *gorm.DB, name string) (*models.QuizMaster, error) {
	frodo, err := models.NewQuizMaster("Frodo Bagins", "1234")
	if err != nil {
		return nil, errors.New("failed to create new user")
	}

	result := db.Create(&frodo)
	if result.Error != nil {
		return nil, errors.New("failed to store new user")
	}

	return frodo, nil
}

func CreateTestQuiz(db *gorm.DB, qmID uuid.UUID, title string) (*models.Quiz, error) {
	q, err := models.NewQuiz(qmID, "XX-XX-XX", title)
	if err != nil {
		return nil, errors.New("failed to create new quiz")
	}

	result := db.Create(&q)
	if result.Error != nil {
		return nil, errors.New("failed to store new user")
	}

	return q, nil
}

func CreateTestSection(db *gorm.DB, qID uuid.UUID, title string) (*models.Section, error) {
	s, err := models.NewSection(qID, title, "")
	if err != nil {
		return nil, errors.New("failed to create new section")
	}

	result := db.Create(&s)
	if result.Error != nil {
		return nil, errors.New("failed to store new section")
	}

	return s, nil
}

func CreateTestQuestion(db *gorm.DB, sID uuid.UUID, question qplugins.QPluginModel) (*models.Question, error) {
	q, err := models.NewQuestion(sID, question)
	if err != nil {
		return nil, errors.New("failed to create new question")
	}

	result := db.Create(&q)
	if result.Error != nil {
		return nil, errors.New("failed to store new question")
	}

	return q, nil
}
