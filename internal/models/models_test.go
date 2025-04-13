package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateTestUser(db *gorm.DB, name string) (*QuizMaster, error) {
	frodo, err := NewQuizMaster("Frodo Bagins", "1234")
	if err != nil {
		return nil, errors.New("failed to create new user")
	}

	result := db.Create(&frodo)
	if result.Error != nil {
		return nil, errors.New("failed to store new user")
	}

	return frodo, nil
}

func CreateTestQuiz(db *gorm.DB, qmID uuid.UUID, title string) (*Quiz, error) {
	q, err := NewQuiz(qmID, "XX-XX-XX", title)
	if err != nil {
		return nil, errors.New("failed to create new quiz")
	}

	result := db.Create(&q)
	if result.Error != nil {
		return nil, errors.New("failed to store new user")
	}

	return q, nil
}

func CreateTestSection(db *gorm.DB, qID uuid.UUID, title string) (*Section, error) {
	s, err := NewSection(qID, title, "")
	if err != nil {
		return nil, errors.New("failed to create new section")
	}

	result := db.Create(&s)
	if result.Error != nil {
		return nil, errors.New("failed to store new section")
	}

	return s, nil
}

type QPluginModelTest struct {
	ID   uuid.UUID
	Text string
}

func NewQPluginModelTest(text string) *QPluginModelTest {
	return &QPluginModelTest{
		ID:   uuid.New(),
		Text: text,
	}
}

func (q *QPluginModelTest) GetType() string {
	return "QPluginModelTest"
}

func (q *QPluginModelTest) GetID() uuid.UUID {
	return q.ID
}
