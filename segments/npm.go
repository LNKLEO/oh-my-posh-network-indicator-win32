package segments

import (
	"github.com/LNKLEO/OMP/properties"
	"github.com/LNKLEO/OMP/runtime"
)

type Npm struct {
	language
}

func (n *Npm) Enabled() bool {
	return n.language.Enabled()
}

func (n *Npm) Template() string {
	return " \ue71e {{.Full}} "
}

func (n *Npm) Init(props properties.Properties, env runtime.Environment) {
	n.language = language{
		env:        env,
		props:      props,
		extensions: []string{"package.json", "package-lock.json"},
		commands: []*cmd{
			{
				executable: "npm",
				args:       []string{"--version"},
				regex:      `(?P<version>((?P<major>[0-9]+).(?P<minor>[0-9]+).(?P<patch>[0-9]+)))`,
			},
		},
		versionURLTemplate: "https://github.com/npm/cli/releases/tag/v{{ .Full }}",
	}
}
