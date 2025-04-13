package sdb

import (
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

		err = Migrator(db)
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
