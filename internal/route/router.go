package route

import "github.com/gin-gonic/gin"

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
}

func (r *Router) RunOrDie(port string) {
	if err := r.g.Run(port); err != nil {
		panic(err)
	}
}
