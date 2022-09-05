package runner

import (
	"os"
)

type hostExecutor struct {
	Out    *os.File
	Tasks  []*task
	Status bool
}

func (executor *hostExecutor) Handle(out *os.File, tasks []*task) {
	executor.Tasks = tasks
	executor.Out = out
}

func (executor *hostExecutor) Run(path string) error {
	for _, task := range executor.Tasks {
		if err := task.Run(path, executor.Out); err != nil {
			return err
		}
	}
	return nil
}
