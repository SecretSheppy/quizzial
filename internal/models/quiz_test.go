package models

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewQuiz(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&QuizMaster{}, &Quiz{})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = teardown(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	frodo, err := NewQuizMaster("Frodo Bagins", "1234")
	if err != nil {
		t.Errorf("failed to create Frodo Bagins: %v", err)
	}

	result := db.Create(&frodo)
	if result.Error != nil {
		t.Error(result.Error)
	}

	t.Run("Create a new quiz", func(t *testing.T) {
		q, err := NewQuiz(frodo.QuizMasterID, "XX-XX-XX", "The Lord of the Rings")
		if err != nil {
			t.Errorf("failed to create a new quiz: %v", err)
		}

		result := db.Create(&q)
		if result.Error != nil {
			t.Error(result.Error)
		}

		db.Preload("Quizzes").First(frodo, frodo.QuizMasterID)
		if len(frodo.Quizzes) == 0 {
			t.Errorf("quiz master Frodo Bagins has no quizzes")
		}

		if frodo.Quizzes[0].Title != "The Lord of the Rings" {
			t.Errorf("quiz created with incorrect title")
		}
	})

	t.Run("Create a new quiz without a quiz master ID", func(t *testing.T) {
		_, err := NewQuiz(uuid.Nil, "XX-XX-XX", "The Lord of the Rings")
		if err == nil {
			t.Errorf("unexpected success creating a new quiz without a quiz master ID")
		}
	})

	t.Run("Create a new quiz with a short ID", func(t *testing.T) {
		_, err := NewQuiz(uuid.Nil, "", "The Lord of the Rings")
		if err == nil {
			t.Errorf("unexpected success creating a new quiz with a short ID")
		}
	})

	t.Run("Create a new quiz without a title", func(t *testing.T) {
		_, err := NewQuiz(uuid.Nil, "XX-XX-XX", "")
		if err == nil {
			t.Errorf("unexpected success creating a new quiz without a title")
		}
	})

	t.Run("Delete a quiz", func(t *testing.T) {
		result := db.Delete(frodo.Quizzes[0])
		if result.Error != nil {
			t.Errorf("failed to delete quiz: %v", result.Error)
		}

		db.Preload("Quizzes").First(frodo, frodo.QuizMasterID)
		if len(frodo.Quizzes) != 0 {
			t.Errorf("quiz master Frodo Bagins has quizzes, but should have none")
		}
	})
}
