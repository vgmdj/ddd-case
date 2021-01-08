package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vgmdj/ddd-case/interfaces/httpfacade/facade"
)

func registerManageRouter(m *gin.RouterGroup) {
	manageFacade := facade.NewManageFacade()

	m.GET("/skuactions", manageFacade.SkuActionsList)

}
