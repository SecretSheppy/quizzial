package sdb

import (
	"github.com/SecretSheppy/quizzial/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	once sync.Once
	db   *gorm.DB
)

func Get() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbname()), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		err = migrate(db)
		if err != nil {
			panic(err)
		}
	})

	return db
}

func dbname() string {
	dbname := os.Getenv("DB_DATABASE_PRODUCTION")
	if os.Getenv("OPERATION_MODE") == "development" {
		dbname = os.Getenv("DB_DATABASE_DEVELOPMENT")
	}
	return dbname
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.SectionResource{},
		&models.QuestionResource{},
		&models.QuizMaster{},
		&models.Quiz{},
		&models.Section{})
}
