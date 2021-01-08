package mapper

import (
	"github.com/vgmdj/ddd-case/domain/manage/entity"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent"
)

// ManageMapper the mapper of manage
type ManageMapper struct {
}

// NewManageMapper return the object of manage mapper
func NewManageMapper() *ManageMapper {
	return &ManageMapper{}
}

// SkuActionTemplateToDO transfer po to do
func (mm *ManageMapper) SkuActionTemplateToDO(data []*ent.SkuAction) []entity.SkuActionTemplate {
	result := make([]entity.SkuActionTemplate, 0)
	for _, c := range data {
		if c == nil {
			continue
		}

		result = append(result, entity.SkuActionTemplate{
			ID:      int64(c.ID),
			AppID:   c.AppID,
			SkuName: c.SkuName,
			Action:  c.Action,
		})
	}

	return result
}
