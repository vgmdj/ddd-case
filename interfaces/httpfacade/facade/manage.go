package facade

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vgmdj/ddd-case/app/service"
	"github.com/vgmdj/ddd-case/infrastructure/api"
	"github.com/vgmdj/ddd-case/interfaces/assembler"
)

// ManageFacade manage api controller
type ManageFacade struct {
	srv       *service.ManageService
	assembler *assembler.ManageAssembler
}

// NewManageFacade return the object of manage facade
func NewManageFacade() *ManageFacade {
	return &ManageFacade{
		srv:       service.NewManageService(),
		assembler: assembler.NewManageAssembler(),
	}
}

// SkuActionsList get the list of sku actions api
func (mf *ManageFacade) SkuActionsList(c *gin.Context) {
	offsetStr, ok := c.GetQuery("offset")
	if !ok {
		c.JSON(http.StatusOK, api.NewAPIReponse(api.CheckErrParamsInvalid, "offset can not be null"))
		return
	}

	limitStr, ok := c.GetQuery("limit")
	if !ok {
		c.JSON(http.StatusOK, api.NewAPIReponse(api.CheckErrParamsInvalid, "limit can not be null"))
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusOK, api.NewAPIReponse(api.CheckErrParamsInvalid, "offset must be number"))
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusOK, api.NewAPIReponse(api.CheckErrParamsInvalid, "limit must be number"))
		return
	}

	// use the timeout setting of downstream
	ctx := context.Background()

	listDo, err := mf.srv.GetSkuActionsList(ctx, offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, api.NewAPIReponse(api.Failed, err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.NewAPIReponse(api.Success, api.Success.Error(), mf.assembler.SkuActionsListToDTO(listDo)))
}
