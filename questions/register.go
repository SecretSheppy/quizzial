package questions

import (
	"errors"
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/SecretSheppy/quizzial/questions/multichoice"
	"gorm.io/gorm"
	"sync"
)

var (
	once                        sync.Once
	allQuestions                map[string]qplugins.QPlugin
	allQPluginModels            map[string]qplugins.QPluginModel
	ErrConflictingQuestionTypes = errors.New("conflicting question types")
)

// RegisteredPlugins is the array of all questions registered in the system. A hardcoded approach was chosen over the
// plugin package due to its lack of support and severe limitations.
var registeredQuestions = []qplugins.QPlugin{
	&multichoice.MultiChoice{},
}

func AllQuestions(DB *gorm.DB) map[string]qplugins.QPlugin {
	once.Do(func() {
		allQuestions = make(map[string]qplugins.QPlugin)
		allQPluginModels = make(map[string]qplugins.QPluginModel)

		for _, plugin := range registeredQuestions {
			plugin.Init(DB)

			models := plugin.GetQPluginModels()
			if err := registerQPluginModels(models); err != nil {
				panic(err)
			}

			if err := plugin.Migrate(); err != nil {

			}

			allQuestions[plugin.Data().Name] = plugin
		}
	})

	return allQuestions
}

func registerQPluginModels(models []qplugins.QPluginModel) error {
	for _, model := range models {
		if allQPluginModels[model.GetType()] != nil {
			return ErrConflictingQuestionTypes
		}

		allQPluginModels[model.GetType()] = model
	}

	return nil
}

func AllQPluginModels() map[string]qplugins.QPluginModel {
	return allQPluginModels
}
