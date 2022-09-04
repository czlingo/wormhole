package executor

import "os"

type Executor interface {
	Exec(path string, cmd *Task) error
	Init(out *os.File, task *Task)
}
