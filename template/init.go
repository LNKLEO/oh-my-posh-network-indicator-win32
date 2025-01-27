package template

import (
	"sync"

	"github.com/LNKLEO/OMP/maps"
	"github.com/LNKLEO/OMP/runtime"
)

const (
	// Errors to show when the template handling fails
	InvalidTemplate   = "invalid template text"
	IncorrectTemplate = "unable to create text based on template"

	globalRef = ".$"
)

var (
	shell          string
	env            runtime.Environment
	knownVariables []string
)

func Init(environment runtime.Environment, vars maps.Simple) {
	env = environment
	shell = env.Shell()

	renderPool = sync.Pool{
		New: func() any {
			return newTextPoolObject()
		},
	}

	knownVariables = []string{
		"Root",
		"PWD",
		"AbsolutePWD",
		"PSWD",
		"Folder",
		"Shell",
		"ShellVersion",
		"UserName",
		"HostName",
		"Code",
		"Env",
		"OS",
		"WSL",
		"PromptCount",
		"Segments",
		"SHLVL",
		"Templates",
		"Var",
		"Data",
		"Jobs",
	}

	if Cache != nil {
		return
	}

	loadCache(vars)
}
