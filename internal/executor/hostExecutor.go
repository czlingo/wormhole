package executor

import (
	"os"
	"os/exec"
)

type HostExecutor struct {
	out    *os.File
	task   *Task
	Status bool
}

func (executor *HostExecutor) Init(out *os.File, task *Task) {
	executor.task = task
	executor.out = out
}

func (executor *HostExecutor) Exec(path string, task *Task) error {
	cmd := exec.Command(task.Cmd, task.Args...)
	cmd.Dir = path
	cmd.Stdout = executor.out

	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
