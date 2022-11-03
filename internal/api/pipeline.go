package api

import (
	"net/http"

	"github.com/czlingo/wormhole/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pipeline struct {
	svc service.Pipeline
}

func NewPipeline(db *gorm.DB) *Pipeline {
	return &Pipeline{
		svc: service.NewPipeline(db),
	}
}

func (p *Pipeline) Save(ctx *gin.Context) {
	req := &service.SavePipelineReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	resp, err := p.svc.Save(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (p *Pipeline) Get(ctx *gin.Context) {
	req := &service.GetPipelineReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	resp, err := p.svc.Get(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (p *Pipeline) Delete(ctx *gin.Context) {
	req := &service.DeletePipelineReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	resp, err := p.svc.Delete(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
