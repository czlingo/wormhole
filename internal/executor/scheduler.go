package executor

import "os"

type Scheduler struct {
}

func Handle(job *Job) error {
	for _, step := range job.Steps {
		executor, err := NewExecutor(step.RunsOn)
		if err != nil {
			return err
		}

		// TODO: change id
		out, err := os.Create(job.Name)
		if err != nil {
			return err
		}

		for _, run := range step.Run {
			task := ParseRun(run)
			// FIXME:
			executor.Handle(out, task)
		}
	}
	return nil
}
