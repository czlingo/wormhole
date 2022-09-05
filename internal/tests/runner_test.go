package tests

import (
	"testing"

	"github.com/czlingo/wormhole/internal/runner"
)

func TestRunner(t *testing.T) {
	err := runner.Handle(&runner.Pipeline{
		Name:    "test",
		Version: "test",
		Env:     nil,
		Steps: []runner.Step{
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
