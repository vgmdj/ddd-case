package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vgmdj/ddd-case/domain/manage/entity"
	"github.com/vgmdj/ddd-case/domain/manage/repository/facade"
	"github.com/vgmdj/ddd-case/infrastructure/repository/mapper"
	"github.com/vgmdj/ddd-case/infrastructure/repository/persistence"
)

// ManageDomainService domain service
type ManageDomainService struct {
	repo   facade.IManage
	mapper *mapper.ManageMapper
}

// NewManageDomainService return the object of manage domain service
func NewManageDomainService() *ManageDomainService {
	return &ManageDomainService{
		repo:   persistence.GlobalRepo.ManagePersistenceRepo(),
		mapper: mapper.NewManageMapper(),
	}
}

// CreateSkuActionTemplate create the template
func (mds *ManageDomainService) CreateSkuActionTemplate() error {

	return nil
}

// DisableSkuActionTemplate disable the template
func (mds *ManageDomainService) DisableSkuActionTemplate() error {

	return nil
}

// GetSkuActionsList get the list of sku actions
func (mds *ManageDomainService) GetSkuActionsList() error {

	return nil
}

// GetSkuActionTemplate get the sku action template
func (mds *ManageDomainService) GetSkuActionTemplate(ctx context.Context, offset, limit int) ([]entity.SkuActionTemplate, error) {
	po, err := mds.repo.GetSkuActionList(ctx, offset, limit)
	if err != nil {
		return nil, errors.WithMessage(err, "get sku action list failed")
	}

	return mds.mapper.SkuActionTemplateToDO(po), nil

}

// GetSkuActionStatusCodes get the sku action status codes
func (mds *ManageDomainService) GetSkuActionStatusCodes() error {

	return nil
}
