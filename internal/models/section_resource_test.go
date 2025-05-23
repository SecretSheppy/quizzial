package models

import (
	"github.com/SecretSheppy/quizzial/internal/sdbtest"
	"testing"
)

func TestNewSectionResource(t *testing.T) {
	db, err := sdbtest.Setup()
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&QuizMaster{}, &Quiz{}, &Section{}, &SectionResource{})
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		err = sdbtest.Teardown(db)
		if err != nil {
			t.Error(err)
		}
	})

	frodo, err := CreateTestUser(db, "Frodo Bagins")
	if err != nil {
		t.Error(err)
	}

	quiz, err := CreateTestQuiz(db, frodo.QuizMasterID, "The Lord of the Rings")
	if err != nil {
		t.Error(err)
	}

	sec, err := CreateTestSection(db, quiz.QuizID, "Section 1")
	if err != nil {
		t.Error(err)
	}

	t.Run("Create new section resource", func(t *testing.T) {
		sr, err := NewSectionResource("/x/y/z", "xyz imag", sec.SectionID)
		if err != nil {
			t.Error(err)
		}

		result := db.Create(sr)
		if result.Error != nil {
			t.Error(result.Error)
		}

		db.Preload("Quizzes.Sections.Cover").First(frodo, frodo.QuizMasterID)
		if frodo.Quizzes[0].Sections[0].Cover.ResourceID != sr.ResourceID {
			t.Error("Cover Resource ID does not match expected ID")
		}
	})

	t.Run("Create new section resource with no path", func(t *testing.T) {
		_, err := NewSectionResource("", "xyz imag", sec.SectionID)
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("Create new section resource with no alt text", func(t *testing.T) {
		_, err := NewSectionResource("/x/y/z", "", sec.SectionID)
		if err == nil {
			t.Error(err)
		}
	})
}
