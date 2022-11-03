package mysql

import (
	"github.com/czlingo/wormhole/internal/model"
	"gorm.io/gorm"
)

type PipelineRepo struct {
	db *gorm.DB
}

func NewPipelineRepo(db *gorm.DB) model.PipelineRepo {
	return &PipelineRepo{
		db: db.Table(model.PipelineTable),
	}
}

func (p *PipelineRepo) Insert(vals ...*model.Pipeline) error {
	return p.db.CreateInBatches(vals, len(vals)).Error
}

func (p *PipelineRepo) Query(ids ...string) ([]*model.Pipeline, error) {
	ret := make([]*model.Pipeline, 0, len(ids))
	err := p.db.Where("id in (?)", ids).Find(ret).Error
	return ret, err
}

// func (p *PipelineRepo) SaveOrUpdate(vals ...*model.Pipeline) error {

// }

func (p *PipelineRepo) Delete(ids ...string) error {
	return p.db.Where("id in (?)", ids).Delete(&model.Pipeline{}).Error
}
