package models

import (
	"github.com/SecretSheppy/quizzial/internal/sdbtest"
	"testing"
)

func TestNewQuizMaster(t *testing.T) {
	db, err := sdbtest.Setup()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&QuizMaster{})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		err = sdbtest.Teardown(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create new quiz master", func(t *testing.T) {
		qm, err := NewQuizMaster("Frodo Bagins", "1234")
		if err != nil {
			t.Errorf("failed to create new quiz master instance: %v", err)
		}

		result := db.Create(qm)
		if result.Error != nil {
			t.Fatal(result.Error)
		}

		var frodo QuizMaster
		db.First(&frodo)

		if frodo.Name != "Frodo Bagins" || frodo.Password != "1234" {
			t.Errorf("New quiz master was not created successfully")
		}
	})

	t.Run("Create quiz master with no name", func(t *testing.T) {
		qm, err := NewQuizMaster("", "1234")
		if err == nil {
			t.Errorf("unexpected success creating quiz master with no name")
		}

		result := db.Create(qm)
		if result.Error == nil {
			t.Fatal(result.Error)
		}
	})

	t.Run("Create quiz master with no password", func(t *testing.T) {
		qm, err := NewQuizMaster("Bilbo Bagins", "")
		if err == nil {
			t.Errorf("unexpected success creating quiz master with no password")
		}

		result := db.Create(qm)
		if result.Error == nil {
			t.Fatal(result.Error)
		}
	})

	t.Run("Delete a quiz master", func(t *testing.T) {
		var qm QuizMaster
		result := db.First(&qm).Where("name = ?", "Frodo Bagins")
		if result.Error != nil {
			t.Fatal(result.Error)
		}

		result = db.Delete(qm)
		if result.Error != nil {
			t.Fatal(result.Error)
		}
	})
}
