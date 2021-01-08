package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// SkuAction holds the schema definition for the SkuAction entity.
type SkuAction struct {
	ent.Schema
}

// Fields of the SkuAction.
func (SkuAction) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_id"),
		field.String("sku_name"),
		field.String("sku_name_cn"),
		field.String("action"),
		field.Int("status"),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SkuAction.
func (SkuAction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("config", SkuConfig.Type).StorageKey(
			edge.Column("sku_action_id"),
		),
		edge.To("template", SkuTemplate.Type).StorageKey(
			edge.Column("sku_action_id"),
		),
		edge.To("status_code", SkuStatusCode.Type).StorageKey(
			edge.Column("sku_action_id"),
		),
	}
}

// Config of the SkuAction
func (SkuAction) Config() ent.Config {
	return ent.Config{
		Table: "sku_action",
	}
}
