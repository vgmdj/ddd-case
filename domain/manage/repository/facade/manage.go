package facade

import (
	"context"

	"github.com/vgmdj/ddd-case/infrastructure/repository/ent"
)

// IManage the interface of manage db
type IManage interface {
	Ping(context.Context) error
	GetSkuActionList(ctx context.Context, offset int, limit int) ([]*ent.SkuAction, error)
}
