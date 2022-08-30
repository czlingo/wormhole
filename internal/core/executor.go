package core

type Executor interface {
	Exec(path string, cmds []string) (string, error)
}
