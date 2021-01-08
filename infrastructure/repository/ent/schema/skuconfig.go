package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// SkuConfig holds the schema definition for the SkuConfig entity.
type SkuConfig struct {
	ent.Schema
}

// Fields of the SkuConfig.
func (SkuConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("config"),
		field.Int("status"),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SkuConfig.
func (SkuConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action", SkuAction.Type).Ref("config").Unique(),
	}
}

// Config of the SkuConfig
func (SkuConfig) Config() ent.Config {
	return ent.Config{
		Table: "sku_config",
	}
}
