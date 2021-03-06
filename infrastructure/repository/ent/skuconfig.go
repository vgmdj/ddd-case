// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skuaction"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skuconfig"
)

// SkuConfig is the model entity for the SkuConfig schema.
type SkuConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Config holds the value of the "config" field.
	Config string `json:"config,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkuConfigQuery when eager-loading is set.
	Edges         SkuConfigEdges `json:"edges"`
	sku_action_id *int
}

// SkuConfigEdges holds the relations/edges for other nodes in the graph.
type SkuConfigEdges struct {
	// Action holds the value of the action edge.
	Action *SkuAction
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ActionOrErr returns the Action value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkuConfigEdges) ActionOrErr() (*SkuAction, error) {
	if e.loadedTypes[0] {
		if e.Action == nil {
			// The edge action was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: skuaction.Label}
		}
		return e.Action, nil
	}
	return nil, &NotLoadedError{edge: "action"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SkuConfig) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // config
		&sql.NullInt64{},  // status
		&sql.NullTime{},   // create_at
		&sql.NullTime{},   // update_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*SkuConfig) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // sku_action_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SkuConfig fields.
func (sc *SkuConfig) assignValues(values ...interface{}) error {
	if m, n := len(values), len(skuconfig.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	sc.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field config", values[0])
	} else if value.Valid {
		sc.Config = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[1])
	} else if value.Valid {
		sc.Status = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_at", values[2])
	} else if value.Valid {
		sc.CreateAt = value.Time
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_at", values[3])
	} else if value.Valid {
		sc.UpdateAt = value.Time
	}
	values = values[4:]
	if len(values) == len(skuconfig.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field sku_action_id", value)
		} else if value.Valid {
			sc.sku_action_id = new(int)
			*sc.sku_action_id = int(value.Int64)
		}
	}
	return nil
}

// QueryAction queries the action edge of the SkuConfig.
func (sc *SkuConfig) QueryAction() *SkuActionQuery {
	return (&SkuConfigClient{config: sc.config}).QueryAction(sc)
}

// Update returns a builder for updating this SkuConfig.
// Note that, you need to call SkuConfig.Unwrap() before calling this method, if this SkuConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SkuConfig) Update() *SkuConfigUpdateOne {
	return (&SkuConfigClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (sc *SkuConfig) Unwrap() *SkuConfig {
	tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SkuConfig is not a transactional entity")
	}
	sc.config.driver = tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SkuConfig) String() string {
	var builder strings.Builder
	builder.WriteString("SkuConfig(")
	builder.WriteString(fmt.Sprintf("id=%v", sc.ID))
	builder.WriteString(", config=")
	builder.WriteString(sc.Config)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sc.Status))
	builder.WriteString(", create_at=")
	builder.WriteString(sc.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(sc.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SkuConfigs is a parsable slice of SkuConfig.
type SkuConfigs []*SkuConfig

func (sc SkuConfigs) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}
