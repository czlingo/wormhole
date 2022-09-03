package core

import "os/exec"

type HostExecutor struct {
}

func (he *HostExecutor) Exec(path string, task *Task) error {
	cmd := exec.Command(task.Cmd, task.Args...)
	cmd.Dir = path
	return cmd.Run()
}
