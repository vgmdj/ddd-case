package assembler

import (
	"github.com/vgmdj/ddd-case/domain/manage/entity"
	"github.com/vgmdj/ddd-case/interfaces/dto"
)

// ManageAssembler the assembler of manage
type ManageAssembler struct {
}

// NewManageAssembler return the object of manage assembler
func NewManageAssembler() *ManageAssembler {
	return &ManageAssembler{}
}

// SkuActionsListToDTO transfer the do to dto
func (ma *ManageAssembler) SkuActionsListToDTO(data []entity.SkuActionTemplate) []dto.SkuActionsListDTO {
	result := make([]dto.SkuActionsListDTO, len(data))

	for i, c := range data {
		result[i] = dto.SkuActionsListDTO{
			ID:      c.ID,
			AppID:   c.AppID,
			SkuName: c.SkuName,
			Action:  c.Action,
		}
	}

	return result
}

// SkuActionsListToDO transfer the dto to do
func (ma *ManageAssembler) SkuActionsListToDO(data []dto.SkuActionsListDTO) []entity.SkuActionTemplate {
	return nil
}
