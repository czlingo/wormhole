package core

import "testing"

func TestHostExec(t *testing.T) {
	he := &HostExecutor{}
	if err := he.Exec("/home/chazhiling/Application", []string{"ls ../"}); err != nil {
		t.Fatal(err)
	}
}
