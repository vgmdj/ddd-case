package service

import (
	"context"

	"github.com/vgmdj/ddd-case/domain/manage/entity"
	"github.com/vgmdj/ddd-case/domain/manage/service"
)

// ManageService the service layer of manage
type ManageService struct {
	srv *service.ManageDomainService
}

// NewManageService return the object of manage service
func NewManageService() *ManageService {
	return &ManageService{
		srv: service.NewManageDomainService(),
	}
}

// GetSkuActionsList get the list of sku actions
func (ms *ManageService) GetSkuActionsList(ctx context.Context, offset int, limit int) ([]entity.SkuActionTemplate, error) {
	return ms.srv.GetSkuActionTemplate(ctx, offset, limit)
}
