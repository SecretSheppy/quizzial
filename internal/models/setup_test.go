package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func teardown(DB *gorm.DB) error {
	var tables []string
	DB.Raw("SELECT name FROM sqlite_master WHERE type='table';").Scan(&tables)

	for _, table := range tables {
		if err := DB.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}

func createTestUser(db *gorm.DB, name string) (*QuizMaster, error) {
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

func createTestQuiz(db *gorm.DB, qmID uuid.UUID, title string) (*Quiz, error) {
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

func createTestSection(db *gorm.DB, qID uuid.UUID, title string) (*Section, error) {
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

type qPluginModelTest struct {
	ID   uuid.UUID
	Text string
}

func newQPluginModelTest(text string) *qPluginModelTest {
	return &qPluginModelTest{
		ID:   uuid.New(),
		Text: text,
	}
}

func (q *qPluginModelTest) GetType() string {
	return "qPluginModelTest"
}

func (q *qPluginModelTest) GetID() uuid.UUID {
	return q.ID
}
