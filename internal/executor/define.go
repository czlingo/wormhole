package executor

const (
	LINUX   = "linux"
	WINDOWS = "windows"
	DOCKER  = "docker"
)

type Job struct {
	Name    string
	Version string
	Env     map[string]string
	Steps   []Step
}

type Step struct {
	Name   string
	Needs  []string
	RunsOn string
	Run    []string
}
