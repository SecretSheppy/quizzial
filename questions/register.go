package questions

import (
	"github.com/SecretSheppy/quizzial/pkg/qplugins"
	"github.com/SecretSheppy/quizzial/questions/multichoice"
	"sync"
)

var (
	once                        sync.Once
	allQuestions                map[string]qplugins.QPlugin
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
	return questions
}
