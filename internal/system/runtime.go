package system

import (
	"github.com/SecretSheppy/quizzial/internal/sdb"
	"github.com/SecretSheppy/quizzial/questions"
	"github.com/joho/godotenv"
)

func Runtime() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	DB := sdb.Get()
	qs := questions.All()

	for _, q := range qs {
		q.Init(DB)

		if err = q.Migrate(); err != nil {
			panic(err)
		}
	}
}
