package runner

import (
	"os"

	"github.com/czlingo/wormhole/internal/errors"
)

type executor interface {
	Run(path string) error
	Handle(out *os.File, task []*task)
}

func newExecutor(runsOn string) (executor, error) {
	switch runsOn {
	case LINUX, WINDOWS:
		return &hostExecutor{}, nil
	case DOCKER:
		// TODO:
	}
	return nil, errors.New(errors.NotSupportRunnerType, runsOn)
}
