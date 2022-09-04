package executor

import (
	"os"

	"github.com/czlingo/wormhole/internal/errors"
)

type Executor interface {
	Exec(path string) error
	Handle(out *os.File, task *Task)
}

func NewExecutor(runsOn string) (Executor, error) {
	switch runsOn {
	case LINUX, WINDOWS:
		return &HostExecutor{}, nil
	case DOCKER:
		// TODO:
	}
	return nil, errors.New(errors.NotSupportRunnerType, runsOn)
}
