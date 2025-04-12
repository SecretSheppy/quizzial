package qplugins

import (
	"gorm.io/gorm"
	"net/http"
)

// QPluginData stores the paths of the local data required by each QPlugin so that it can be registered and used
// by the main system.
type QPluginData struct {
	Name string
}

// QPlugin is a question plugin. All questions must implement the QPlugin interface. Any extra question plugins
// created must be added to the RegisteredQuestions array in the `qplugins/register.go` file.
type QPlugin interface {

	// Data returns all known data about the QPlugin
	Data() *QPluginData

	// Init initialises the QPlugin with the database variable so that it can be stored in the plugin struct and does
	// not have to be a parameter in every function that uses the database.
	Init(db *gorm.DB)

	// Migrate migrates the gorm.Model(s) in to the database if they are not already present.
	Migrate() error

	// SaveQuestionHandler saves the current question configuration into the database.
	SaveQuestionHandler(w http.ResponseWriter, r *http.Request)

	// VerifyQuestion verifies whether the mandatory parts of the question are completed. This must return false if
	// core parts of the question are missing or incomplete as this function is run when the user attempts to open
	// a quiz for public participation.
	VerifyQuestion(ID int64) bool

	// SaveAnswersHandler saves the question answers for the current user into the database using the locally defined
	// models.
	SaveAnswersHandler(w http.ResponseWriter, r *http.Request)
}
