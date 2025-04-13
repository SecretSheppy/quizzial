package models

import (
	"github.com/SecretSheppy/quizzial/internal/sdbtest"
	"github.com/google/uuid"
	"testing"
)

func TestNewQuestion(t *testing.T) {
	db, err := sdbtest.Setup()
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&QuizMaster{}, &Quiz{}, &Section{}, &Question{})
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

	t.Run("Create new question", func(t *testing.T) {
		question := NewQPluginModelTest("hi there")

		q, err := NewQuestion(sec.SectionID, question)
		if err != nil {
			t.Error(err)
		}

		result := db.Create(q)
		if result.Error != nil {
			t.Error(result.Error)
		}

		db.Preload("Quizzes.Sections.Questions").First(frodo, frodo.QuizMasterID)
		if frodo.Quizzes[0].Sections[0].Questions[0].QuestionableType != question.GetType() {
			t.Error("Questionable type is not correct")
		}
	})

	t.Run("Create new question without sID", func(t *testing.T) {
		question := NewQPluginModelTest("hi there")

		_, err := NewQuestion(uuid.Nil, question)
		if err == nil {
			t.Error("Question creation did not fail")
		}
	})

	t.Run("Create new question without QPluginModel", func(t *testing.T) {
		_, err := NewQuestion(sec.SectionID, nil)
		if err == nil {
			t.Error("Question creation did not fail")
		}
	})
}
