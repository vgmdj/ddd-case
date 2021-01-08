package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vgmdj/ddd-case/interfaces/httpfacade/middleware"
)

// Register resigter gin router
func Register(g *gin.Engine) {
	g.HEAD("ping", middleware.PingHealth)
	g.NoRoute(middleware.NotFindRouter)

	g.Use(
		middleware.GinQscLogger(),
		middleware.GinQscRecovery(),
	)

	m := g.Group("/mock")
	registerMockRouter(m)

	mg := g.Group("/manage")
	registerManageRouter(mg)

	f := g.Group("/fork")
	registerForkRouter(f)

}
