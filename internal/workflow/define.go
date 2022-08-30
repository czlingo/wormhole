package workflow

type Action struct {
	Version string
	Name    string
	Env     map[string]string
	Steps   []Step
}

type Step struct {
	Name   string
	Needs  []string
	RunsOn string
	Run    []string
}
