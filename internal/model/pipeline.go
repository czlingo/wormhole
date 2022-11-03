package model

const (
	PipelineTable = "pipeline"
)

type Pipeline struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Content    []byte `json:"content"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type PipelineRepo interface {
	Insert(...*Pipeline) error
	Query(ids ...string) ([]*Pipeline, error)
	// SaveOrUpdate(...*Pipeline) error
	Delete(ids ...string) error
}
