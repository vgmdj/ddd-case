package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// SkuTemplate holds the schema definition for the SkuTemplate entity.
type SkuTemplate struct {
	ent.Schema
}

// Fields of the SkuTemplate.
func (SkuTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("request"),
		field.String("response"),
		field.Int("status"),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SkuTemplate.
func (SkuTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action", SkuAction.Type).Ref("template").Unique(),
	}
}

// Config of the SkuTemplate
func (SkuTemplate) Config() ent.Config {
	return ent.Config{
		Table: "sku_template",
	}
}
