package models

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewQuestion(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&QuizMaster{}, &Quiz{}, &Section{}, &Question{})
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

	sec, err := createTestSection(db, quiz.QuizID, "Section 1")
	if err != nil {
		t.Error(err)
	}

	t.Run("Create new question", func(t *testing.T) {
		question := newQPluginModelTest("hi there")

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
		question := newQPluginModelTest("hi there")

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
