package entity

import "github.com/vgmdj/ddd-case/domain/manage/entity/valueobject"

// SkuActionTemplate the whole config template fo sku action
type SkuActionTemplate struct {
	ID       int64
	AppID    string
	SkuName  string
	Action   string
	Config   valueobject.Config
	Request  valueobject.Request
	Response valueobject.Response
}
