package tests

import (
	"testing"

	"github.com/czlingo/wormhole/internal/define"
	"github.com/czlingo/wormhole/internal/runner"
)

func TestRunner(t *testing.T) {
	err := runner.Handle(&define.Pipeline{
		Name:    "test",
		Version: "test",
		Env:     nil,
		Steps: []define.Step{
			{
				Name:   "clone",
				RunsOn: "linux",
				Run: []string{
					"touch test",
					"ls",
					"rm test",
				},
			},
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
}
