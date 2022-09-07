package define

import (
	"gopkg.in/yaml.v3"
)

type Pipeline struct {
	Name    string            `yaml:"name"`
	Version string            `yaml:"version"`
	Env     map[string]string `yaml:"env"`
	Steps   []Step            `yaml:"steps"`
}

type Step struct {
	Name   string   `yaml:"name"`
	Needs  []string `yaml:"needs"`
	RunsOn string   `yaml:"runsOn"`
	Run    []string `yaml:"run"`
}

func ParsePipeline(d []byte) (*Pipeline, error) {
	pipe := &Pipeline{}
	err := yaml.Unmarshal(d, pipe)
	return pipe, err
}
