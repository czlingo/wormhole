package executor

import "strings"

type Task struct {
	ID   string
	Cmd  string
	Args []string
}

func ParseRun(run string) *Task {
	if run == "" {
		return nil
	}

	words := strings.Split(run, " ")
	return &Task{
		Cmd:  words[0],
		Args: words[1:],
	}
}
