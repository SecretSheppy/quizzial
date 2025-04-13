package system

import (
	"fmt"
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
	qs := questions.AllQuestions(DB)
	qt := questions.AllQPluginModels()

	fmt.Println("All Questions Initialized")
	fmt.Println(qs)
	fmt.Println(qt)
}
