package core

import (
	"os/exec"
	"strings"
)

type HostExecutor struct {
}

func (he *HostExecutor) Exec(path string, cmds []string) error {
	for _, cmd := range cmds {
		words := strings.Split(cmd, " ")
		ec := exec.Command(words[0], words[1:]...)
		ec.Dir = path
		if err := ec.Run(); err != nil {
			return err
		}
	}
	return nil
}
