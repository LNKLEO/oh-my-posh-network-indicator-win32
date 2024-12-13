package segments

import (
	"github.com/LNKLEO/OMP/properties"
	"github.com/LNKLEO/OMP/runtime"
)

type Deno struct {
	language
}

func (d *Deno) Template() string {
	return languageTemplate
}

func (d *Deno) Init(props properties.Properties, env runtime.Environment) {
	d.language = language{
		env:        env,
		props:      props,
		extensions: []string{"*.js", "*.ts", "deno.json"},
		commands: []*cmd{
			{
				executable: "deno",
				args:       []string{"--version"},
				regex:      `(?:(?P<version>((?P<major>[0-9]+).(?P<minor>[0-9]+).(?P<patch>[0-9]+))))`,
			},
		},
		versionURLTemplate: "https://github.com/denoland/deno/releases/tag/v{{.Full}}",
	}
}

func (d *Deno) Enabled() bool {
	return d.language.Enabled()
}
