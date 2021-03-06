// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skuaction"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skuconfig"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skustatuscode"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skutemplate"
)

// SkuActionCreate is the builder for creating a SkuAction entity.
type SkuActionCreate struct {
	config
	mutation *SkuActionMutation
	hooks    []Hook
}

// SetAppID sets the app_id field.
func (sac *SkuActionCreate) SetAppID(s string) *SkuActionCreate {
	sac.mutation.SetAppID(s)
	return sac
}

// SetSkuName sets the sku_name field.
func (sac *SkuActionCreate) SetSkuName(s string) *SkuActionCreate {
	sac.mutation.SetSkuName(s)
	return sac
}

// SetSkuNameCn sets the sku_name_cn field.
func (sac *SkuActionCreate) SetSkuNameCn(s string) *SkuActionCreate {
	sac.mutation.SetSkuNameCn(s)
	return sac
}

// SetAction sets the action field.
func (sac *SkuActionCreate) SetAction(s string) *SkuActionCreate {
	sac.mutation.SetAction(s)
	return sac
}

// SetStatus sets the status field.
func (sac *SkuActionCreate) SetStatus(i int) *SkuActionCreate {
	sac.mutation.SetStatus(i)
	return sac
}

// SetCreateAt sets the create_at field.
func (sac *SkuActionCreate) SetCreateAt(t time.Time) *SkuActionCreate {
	sac.mutation.SetCreateAt(t)
	return sac
}

// SetNillableCreateAt sets the create_at field if the given value is not nil.
func (sac *SkuActionCreate) SetNillableCreateAt(t *time.Time) *SkuActionCreate {
	if t != nil {
		sac.SetCreateAt(*t)
	}
	return sac
}

// SetUpdateAt sets the update_at field.
func (sac *SkuActionCreate) SetUpdateAt(t time.Time) *SkuActionCreate {
	sac.mutation.SetUpdateAt(t)
	return sac
}

// SetNillableUpdateAt sets the update_at field if the given value is not nil.
func (sac *SkuActionCreate) SetNillableUpdateAt(t *time.Time) *SkuActionCreate {
	if t != nil {
		sac.SetUpdateAt(*t)
	}
	return sac
}

// AddConfigIDs adds the config edge to SkuConfig by ids.
func (sac *SkuActionCreate) AddConfigIDs(ids ...int) *SkuActionCreate {
	sac.mutation.AddConfigIDs(ids...)
	return sac
}

// AddConfig adds the config edges to SkuConfig.
func (sac *SkuActionCreate) AddConfig(s ...*SkuConfig) *SkuActionCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sac.AddConfigIDs(ids...)
}

// AddTemplateIDs adds the template edge to SkuTemplate by ids.
func (sac *SkuActionCreate) AddTemplateIDs(ids ...int) *SkuActionCreate {
	sac.mutation.AddTemplateIDs(ids...)
	return sac
}

// AddTemplate adds the template edges to SkuTemplate.
func (sac *SkuActionCreate) AddTemplate(s ...*SkuTemplate) *SkuActionCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sac.AddTemplateIDs(ids...)
}

// AddStatusCodeIDs adds the status_code edge to SkuStatusCode by ids.
func (sac *SkuActionCreate) AddStatusCodeIDs(ids ...int) *SkuActionCreate {
	sac.mutation.AddStatusCodeIDs(ids...)
	return sac
}

// AddStatusCode adds the status_code edges to SkuStatusCode.
func (sac *SkuActionCreate) AddStatusCode(s ...*SkuStatusCode) *SkuActionCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sac.AddStatusCodeIDs(ids...)
}

// Mutation returns the SkuActionMutation object of the builder.
func (sac *SkuActionCreate) Mutation() *SkuActionMutation {
	return sac.mutation
}

// Save creates the SkuAction in the database.
func (sac *SkuActionCreate) Save(ctx context.Context) (*SkuAction, error) {
	var (
		err  error
		node *SkuAction
	)
	sac.defaults()
	if len(sac.hooks) == 0 {
		if err = sac.check(); err != nil {
			return nil, err
		}
		node, err = sac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkuActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sac.check(); err != nil {
				return nil, err
			}
			sac.mutation = mutation
			node, err = sac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sac.hooks) - 1; i >= 0; i-- {
			mut = sac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sac *SkuActionCreate) SaveX(ctx context.Context) *SkuAction {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sac *SkuActionCreate) defaults() {
	if _, ok := sac.mutation.CreateAt(); !ok {
		v := skuaction.DefaultCreateAt()
		sac.mutation.SetCreateAt(v)
	}
	if _, ok := sac.mutation.UpdateAt(); !ok {
		v := skuaction.DefaultUpdateAt()
		sac.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *SkuActionCreate) check() error {
	if _, ok := sac.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New("ent: missing required field \"app_id\"")}
	}
	if _, ok := sac.mutation.SkuName(); !ok {
		return &ValidationError{Name: "sku_name", err: errors.New("ent: missing required field \"sku_name\"")}
	}
	if _, ok := sac.mutation.SkuNameCn(); !ok {
		return &ValidationError{Name: "sku_name_cn", err: errors.New("ent: missing required field \"sku_name_cn\"")}
	}
	if _, ok := sac.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New("ent: missing required field \"action\"")}
	}
	if _, ok := sac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := sac.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New("ent: missing required field \"create_at\"")}
	}
	if _, ok := sac.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New("ent: missing required field \"update_at\"")}
	}
	return nil
}

func (sac *SkuActionCreate) sqlSave(ctx context.Context) (*SkuAction, error) {
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sac *SkuActionCreate) createSpec() (*SkuAction, *sqlgraph.CreateSpec) {
	var (
		_node = &SkuAction{config: sac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: skuaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skuaction.FieldID,
			},
		}
	)
	if value, ok := sac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: skuaction.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := sac.mutation.SkuName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: skuaction.FieldSkuName,
		})
		_node.SkuName = value
	}
	if value, ok := sac.mutation.SkuNameCn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: skuaction.FieldSkuNameCn,
		})
		_node.SkuNameCn = value
	}
	if value, ok := sac.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: skuaction.FieldAction,
		})
		_node.Action = value
	}
	if value, ok := sac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: skuaction.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sac.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: skuaction.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := sac.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: skuaction.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if nodes := sac.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skuaction.ConfigTable,
			Columns: []string{skuaction.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skuconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skuaction.TemplateTable,
			Columns: []string{skuaction.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skutemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.StatusCodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skuaction.StatusCodeTable,
			Columns: []string{skuaction.StatusCodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skustatuscode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SkuActionCreateBulk is the builder for creating a bulk of SkuAction entities.
type SkuActionCreateBulk struct {
	config
	builders []*SkuActionCreate
}

// Save creates the SkuAction entities in the database.
func (sacb *SkuActionCreateBulk) Save(ctx context.Context) ([]*SkuAction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*SkuAction, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkuActionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (sacb *SkuActionCreateBulk) SaveX(ctx context.Context) []*SkuAction {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
