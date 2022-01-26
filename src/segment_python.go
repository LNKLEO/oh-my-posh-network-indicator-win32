package main

import (
	"oh-my-posh/environment"
	"oh-my-posh/properties"
)

type python struct {
	language

	Venv string
}

const (
	// FetchVirtualEnv fetches the virtual env
	FetchVirtualEnv properties.Property = "fetch_virtual_env"
)

func (p *python) template() string {
	return languageTemplate
}

func (p *python) init(props properties.Properties, env environment.Environment) {
	p.language = language{
		env:         env,
		props:       props,
		extensions:  []string{"*.py", "*.ipynb", "pyproject.toml", "venv.bak", "venv", ".venv"},
		loadContext: p.loadContext,
		inContext:   p.inContext,
		commands: []*cmd{
			{
				executable: "python",
				args:       []string{"--version"},
				regex:      `(?:Python (?P<version>((?P<major>[0-9]+).(?P<minor>[0-9]+).(?P<patch>[0-9]+))))`,
			},
			{
				executable: "python3",
				args:       []string{"--version"},
				regex:      `(?:Python (?P<version>((?P<major>[0-9]+).(?P<minor>[0-9]+).(?P<patch>[0-9]+))))`,
			},
		},
		versionURLTemplate: "[%s](https://www.python.org/downloads/release/python-%s%s%s/)",
		displayMode:        props.GetString(DisplayMode, DisplayModeEnvironment),
		homeEnabled:        props.GetBool(HomeEnabled, true),
	}
}

func (p *python) enabled() bool {
	return p.language.enabled()
}

func (p *python) loadContext() {
	if !p.language.props.GetBool(FetchVirtualEnv, true) {
		return
	}
	venvVars := []string{
		"VIRTUAL_ENV",
		"CONDA_ENV_PATH",
		"CONDA_DEFAULT_ENV",
		"PYENV_VERSION",
	}
	var venv string
	for _, venvVar := range venvVars {
		venv = p.language.env.Getenv(venvVar)
		name := environment.Base(p.language.env, venv)
		if p.canUseVenvName(name) {
			p.Venv = name
			break
		}
	}
}

func (p *python) inContext() bool {
	return p.Venv != ""
}

func (p *python) canUseVenvName(name string) bool {
	if name == "" || name == "." {
		return false
	}
	if p.language.props.GetBool(properties.DisplayDefault, true) {
		return true
	}
	invalidNames := [2]string{"system", "base"}
	for _, a := range invalidNames {
		if a == name {
			return false
		}
	}
	return true
}
