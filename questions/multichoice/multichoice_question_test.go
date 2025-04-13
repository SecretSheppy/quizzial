package multichoice

import (
	"github.com/SecretSheppy/quizzial/internal/modelstest"
	"github.com/SecretSheppy/quizzial/internal/sdbtest"
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/google/uuid"
	"testing"
)

func TestNewMultiChoiceQuestion(t *testing.T) {
	db, err := sdbtest.Setup()
	if err != nil {
		t.Error(err)
	}

	err = modelstest.MigrateAll(db)
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&Option{}, &MultiChoiceQuestion{})
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		err = sdbtest.Teardown(db)
		if err != nil {
			t.Error(err)
		}
	})

	mp := make(map[string]qplugins.QPluginModel)
	mp["MultiChoiceQuestion"] = &MultiChoiceQuestion{}

	frodo, err := modelstest.CreateTestUser(db, "Frodo Bagins")
	if err != nil {
		t.Error(err)
	}

	quiz, err := modelstest.CreateTestQuiz(db, frodo.QuizMasterID, "The Lord of the Rings")
	if err != nil {
		t.Error(err)
	}

	sec, err := modelstest.CreateTestSection(db, quiz.QuizID, "Section 1")
	if err != nil {
		t.Error(err)
	}

	t.Run("Create a new multichoice question", func(t *testing.T) {
		mcq, err := NewMultiChoiceQuestion("Where is the ring?")
		if err != nil {
			t.Error(err)
		}

		result := db.Create(mcq)
		if result.Error != nil {
			t.Error(result.Error)
		}

		q, err := modelstest.CreateTestQuestion(db, sec.SectionID, mcq)
		if err != nil {
			t.Error(err)
		}

		db.Preload("Quizzes.Sections.Questions").First(frodo, frodo.QuizMasterID)
		if frodo.Quizzes[0].Sections[0].Questions[0].QuestionID != q.QuestionID {
			t.Error("Question does not match")
		}

		model, err := q.GetQuestionModel(db, mp)
		if err != nil {
			t.Error(err)
		}

		if model.(*MultiChoiceQuestion).GetType() != "MultiChoiceQuestion" {
			t.Error("Question does not match")
		}
	})

	t.Run("Create a new multichoice question with five options", func(t *testing.T) {
		mcq, err := NewMultiChoiceQuestion("Where is the ring?")
		if err != nil {
			t.Error(err)
		}

		mcq.Options = []Option{
			{
				OptionID:              uuid.New(),
				MultiChoiceQuestionID: mcq.MultiChoiceQuestionID,
				Text:                  "option 1",
				Correct:               true,
			},
			{
				OptionID:              uuid.New(),
				MultiChoiceQuestionID: mcq.MultiChoiceQuestionID,
				Text:                  "option 2",
				Correct:               false,
			},
			{
				OptionID:              uuid.New(),
				MultiChoiceQuestionID: mcq.MultiChoiceQuestionID,
				Text:                  "option 3",
				Correct:               true,
			},
			{
				OptionID:              uuid.New(),
				MultiChoiceQuestionID: mcq.MultiChoiceQuestionID,
				Text:                  "option 4",
				Correct:               false,
			},
			{
				OptionID:              uuid.New(),
				MultiChoiceQuestionID: mcq.MultiChoiceQuestionID,
				Text:                  "option 5",
				Correct:               false,
			},
		}

		result := db.Create(mcq)
		if result.Error != nil {
			t.Error(result.Error)
		}

		q, err := modelstest.CreateTestQuestion(db, sec.SectionID, mcq)
		if err != nil {
			t.Error(err)
		}

		db.Preload("Quizzes.Sections.Questions").First(frodo, frodo.QuizMasterID)

		question := frodo.Quizzes[0].Sections[0].Questions[1]
		if question.QuestionID != q.QuestionID {
			t.Error("Question does not match")
		}

		model, err := q.GetQuestionModel(db, mp)
		if err != nil {
			t.Error(err)
		}

		val, ok := model.(*MultiChoiceQuestion)
		if !ok {
			t.Error("Question is not MultiChoiceQuestion")
		}

		if len(val.Options) != 5 {
			t.Error("Options length does not match")
		}
	})
}
