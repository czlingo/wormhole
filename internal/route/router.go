package route

import (
	"strings"

	"github.com/czlingo/wormhole/internal/api"
	"github.com/gin-gonic/gin"
)

type Router struct {
	g *gin.Engine
}

func New() *Router {
	r := &Router{
		g: gin.Default(),
	}
	initRoute(r)
	return r
}

func initRoute(r *Router) {
	// TODO:
	apiBase := r.g.Group("/api")

	{
		apiBase.GET("/ping", api.Ping)
	}

	pipeline := api.NewPipeline()
	pipelineGroup := apiBase.Group("/pipeline")
	{
		pipelineGroup.POST("", pipeline.Save)
		pipelineGroup.DELETE("/:id", pipeline.Delete)
		pipelineGroup.GET("/:id", pipeline.Get)
	}
}

func (r *Router) RunOrDie(port string) {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	if err := r.g.Run(port); err != nil {
		panic(err)
	}
}
