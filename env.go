package env

import (
	"fmt"
	"os"
)

var envs []Env

type Env struct {
	Name     string
	Describe string
}

func New(name string, desc ...string) Env {
	if len(desc) > 0 {
		return Env{Name: name, Describe: desc[0]}
	}
	return Env{Name: name}
}

func Add(e Env) {
	envs = append(envs, e)
}

func PrintDefaults() {
	var s string
	for _, v := range envs {
		s += fmt.Sprintf("  - %s (current \"%s\")\n", v.Name, os.Getenv(v.Name))
		if len(v.Describe) > 0 {
			s += fmt.Sprintf("    \t%s\n", v.Describe)
		}
	}
	fmt.Printf(s)
}
