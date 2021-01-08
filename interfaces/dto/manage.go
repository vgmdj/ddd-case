package dto

// SkuActionsListDTO the dto of sku actions list
type SkuActionsListDTO struct {
	ID      int64  `json:"id"`
	AppID   string `json:"app_id"`
	SkuName string `json:"sku_name"`
	Action  string `json:"action"`
}
