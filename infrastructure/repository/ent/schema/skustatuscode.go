package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// SkuStatusCode holds the schema definition for the SkuStatusCode entity.
type SkuStatusCode struct {
	ent.Schema
}

// Fields of the SkuStatusCode.
func (SkuStatusCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("insurance_code"),
		field.String("reflect_code"),
		field.Int("status"),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SkuStatusCode.
func (SkuStatusCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action", SkuAction.Type).Ref("status_code").Unique(),
	}
}

// Config of the SkuStatusCode
func (SkuStatusCode) Config() ent.Config {
	return ent.Config{
		Table: "sku_status_code",
	}
}
