package multichoice

import (
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"gorm.io/gorm"
	"net/http"
)

var config = &qplugins.QPluginData{
	Name: "multichoice",
}

type MultiChoice struct {
	DB *gorm.DB
}

func (m *MultiChoice) Data() *qplugins.QPluginData {
	return config
}

func (m *MultiChoice) Init(db *gorm.DB) {
	m.DB = db
}

func (m *MultiChoice) Migrate() {

}

func (m *MultiChoice) SaveQuestionHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *MultiChoice) VerifyQuestion(ID int64) bool {
	return true
}

func (m *MultiChoice) SaveAnswersHandler(w http.ResponseWriter, r *http.Request) {

}
