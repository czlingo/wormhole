package executor

import (
	"os"
	"os/exec"
)

type HostExecutor struct {
	Out    *os.File
	Task   *Task
	Status bool
}

func (executor *HostExecutor) Init(out *os.File, task *Task) {
	executor.Task = task
	executor.Out = out
}

func (executor *HostExecutor) Exec(path string, task *Task) error {
	cmd := exec.Command(task.Cmd, task.Args...)
	cmd.Dir = path
	cmd.Stdout = executor.Out

	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
