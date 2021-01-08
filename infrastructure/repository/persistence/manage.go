package persistence

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vgmdj/ddd-case/domain/manage/repository/facade"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent"
)

var _ facade.IManage = (*ManageRepository)(nil)

// ManageRepository imp the interface of domain.IManage
type ManageRepository struct {
	cli *ent.Client
}

// NewManageRepository return the object of manage repo
func NewManageRepository(cli *ent.Client) *ManageRepository {
	return &ManageRepository{
		cli: cli,
	}
}

// Ping check the manage conn is ok or not
func (mr *ManageRepository) Ping(ctx context.Context) error {
	ok, err := mr.cli.SkuAction.Query().Exist(ctx)
	if err != nil {
		return errors.WithMessage(err, "sku action get failed")
	}

	if !ok {
		return errors.New("sku action is not exist")
	}

	return nil
}

// GetSkuActionList get the sku action list
func (mr *ManageRepository) GetSkuActionList(ctx context.Context, offset, limit int) ([]*ent.SkuAction, error) {
	return mr.cli.SkuAction.Query().Offset(offset).Limit(limit).All(ctx)
}
