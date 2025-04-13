package models

import (
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
