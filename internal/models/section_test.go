package models

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewSection(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&QuizMaster{}, &Quiz{}, &Section{})
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		err = teardown(db)
		if err != nil {
			t.Error(err)
		}
	})

	frodo, err := createTestUser(db, "Frodo Bagins")
	if err != nil {
		t.Error(err)
	}

	quiz, err := createTestQuiz(db, frodo.QuizMasterID, "The Lord of the Rings")
	if err != nil {
		t.Error(err)
	}

	t.Run("Create a new section", func(t *testing.T) {
		s, err := NewSection(quiz.QuizID, "Section 1", "bla bla bla...")
		if err != nil {
			t.Error(err)
		}

		result := db.Create(&s)
		if result.Error != nil {
			t.Error(result.Error)
		}

		db.Preload("Quizzes.Sections").First(frodo, frodo.QuizMasterID)
		if frodo.Quizzes == nil || len(frodo.Quizzes) == 0 {
			t.Error("quiz master Frodo Bagins has no quizzes")
		}

		if frodo.Quizzes[0].Sections == nil || len(frodo.Quizzes[0].Sections) == 0 {
			t.Error("quiz master Frodo Bagins has no sections")
		}

		if frodo.Quizzes[0].Sections[0].Title != "Section 1" {
			t.Error("Section 1 created incorrectly")
		}
	})

	t.Run("Create a new section without a quiz id", func(t *testing.T) {
		_, err := NewSection(uuid.Nil, "Section 1", "bla bla...")
		if err == nil {
			t.Error("expected error when creating a new section without a quiz id")
		}
	})

	t.Run("Create a new section without a title", func(t *testing.T) {
		_, err := NewSection(quiz.QuizMasterID, "", "bla bla...")
		if err == nil {
			t.Error("expected error when creating a new section without a title")
		}
	})

	t.Run("Create a new section without a description", func(t *testing.T) {
		_, err := NewSection(quiz.QuizMasterID, "Section 1", "bla bla...")
		if err != nil {
			t.Errorf("unexpected error when creating a new section without a description: %v", err)
		}
	})

	t.Run("Delete a section", func(t *testing.T) {
		result := db.Delete(frodo.Quizzes[0].Sections[0])
		if result.Error != nil {
			t.Error(result.Error)
		}

		db.Preload("Quizzes.Sections").First(frodo, frodo.QuizMasterID)
		if len(frodo.Quizzes[0].Sections) != 0 {
			t.Errorf("quiz master Frodo Bagins has sections in quiz 1 despite section being deleted")
		}
	})
}
