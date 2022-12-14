package runner

import (
	"os"

	"github.com/czlingo/wormhole/config"
	"github.com/czlingo/wormhole/internal/define"
)

func Handle(pipe *define.Pipeline) error {
	// TODO: change id & change path
	out, err := os.Create(pipe.Name)
	if err != nil {
		return err
	}

	for _, step := range pipe.Steps {
		executor, err := newExecutor(step.RunsOn)
		if err != nil {
			return err
		}

		tasks := make([]*task, 0, 8)
		for _, run := range step.Run {
			tasks = append(tasks, parseRun(run))
		}

		executor.Handle(out, tasks)
		// TODO: change path
		if err := executor.Run(config.DATA_PATH); err != nil {
			return err
		}
	}
	return nil
}
