package runner

import (
	"os"
	"os/exec"
	"strings"
)

type task struct {
	ID   string
	Cmd  string
	Args []string
}

func parseRun(run string) *task {
	if run == "" {
		return nil
	}

	// FIXME: "hello world"
	words := strings.Split(run, " ")
	return &task{
		Cmd:  words[0],
		Args: words[1:],
	}
}

func (t *task) Run(path string, out *os.File) error {
	cmd := exec.Command(t.Cmd, t.Args...)
	cmd.Dir = path
	cmd.Stdout = out
	cmd.Stderr = out

	return cmd.Run()
}
