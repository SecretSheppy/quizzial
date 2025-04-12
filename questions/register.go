package questions

import (
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/SecretSheppy/quizzial/questions/multichoice"
)

// RegisteredPlugins is the array of all questions registered in the system. A hardcoded approach was chosen over the
// plugin package due to its lack of support and severe limitations.
var registeredQuestions = []qplugins.QPlugin{
	&multichoice.MultiChoice{},
}

func All() map[string]qplugins.QPlugin {
	var questions map[string]qplugins.QPlugin
	for _, plugin := range registeredQuestions {
		questions[plugin.Data().Name] = plugin
	}
	return questions
}
