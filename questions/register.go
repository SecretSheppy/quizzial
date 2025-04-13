package questions

import (
	"errors"
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/SecretSheppy/quizzial/questions/multichoice"
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

func AllQuestions() map[string]qplugins.QPlugin {
	once.Do(func() {
		allQuestions = make(map[string]qplugins.QPlugin)
		for _, plugin := range registeredQuestions {
			allQuestions[plugin.Data().Name] = plugin
		}
	})

	return allQuestions
}

func RegisterQPluginModel(models ...qplugins.QPluginModel) error {
	once.Do(func() {
		allQPluginModels = make(map[string]qplugins.QPluginModel)
	})

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
