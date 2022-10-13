package service

import (
	"time"

	"github.com/czlingo/wormhole/internal/define"
	"github.com/czlingo/wormhole/internal/model"
)

type Pipeline interface {
	Save(*SavePipelineReq) (*SavePipelineResp, error)
	Get(*GetPipelineReq) (*GetPipelineResp, error)
	List(*ListPipelineReq) (*ListPipelineResp, error)
	Delete(*DeletePipelineReq) (*DeletePipelineResp, error)
}

type pipeline struct {
	repo model.PipelineRepo
}

func NewPipeline() Pipeline {
	// TODO: init
	return &pipeline{}
}

type SavePipelineReq struct {
	Content []byte `json:"content"`
}

type SavePipelineResp struct {
	Name string `json:"name"`
}

func (p *pipeline) Save(req *SavePipelineReq) (*SavePipelineResp, error) {
	curTime := time.Now().UTC().Unix()

	pipeline, err := define.ParsePipeline(req.Content)
	if err != nil {
		return nil, err
	}

	e := &model.Pipeline{
		ID:         "",
		Name:       pipeline.Name,
		Content:    req.Content,
		CreateTime: curTime,
		UpdateTime: curTime,
	}

	if err := p.repo.Insert(e); err != nil {
		return nil, err
	}

	return &SavePipelineResp{e.Name}, nil
}

type GetPipelineReq struct {
	ID string `uri:"id"`
}

type GetPipelineResp = model.Pipeline

func (p *pipeline) Get(req *GetPipelineReq) (*GetPipelineResp, error) {
	pipes, err := p.repo.Query(req.ID)
	if err != nil {
		return nil, err
	}
	return pipes[0], nil
}

type ListPipelineReq struct {
}

type ListPipelineResp = []*model.Pipeline

func (p *pipeline) List(req *ListPipelineReq) (*ListPipelineResp, error) {
	panic("not implement")
}

type DeletePipelineReq struct {
	ID string `uri:"id"`
}

type DeletePipelineResp struct {
}

func (p *pipeline) Delete(req *DeletePipelineReq) (*DeletePipelineResp, error) {
	if err := p.repo.Delete(req.ID); err != nil {
		return nil, err
	}
	return &DeletePipelineResp{}, nil
}
