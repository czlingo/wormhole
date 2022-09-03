package core

type Executor interface {
	Exec(path string, cmd *Task) error
}
