package executor

import "os"

type Job struct {
	Tasks []*Task
	Out   *os.File
}
