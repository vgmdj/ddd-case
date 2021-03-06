// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/predicate"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skuconfig"
)

// SkuConfigDelete is the builder for deleting a SkuConfig entity.
type SkuConfigDelete struct {
	config
	hooks    []Hook
	mutation *SkuConfigMutation
}

// Where adds a new predicate to the delete builder.
func (scd *SkuConfigDelete) Where(ps ...predicate.SkuConfig) *SkuConfigDelete {
	scd.mutation.predicates = append(scd.mutation.predicates, ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SkuConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(scd.hooks) == 0 {
		affected, err = scd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkuConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			scd.mutation = mutation
			affected, err = scd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(scd.hooks) - 1; i >= 0; i-- {
			mut = scd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SkuConfigDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SkuConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: skuconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skuconfig.FieldID,
			},
		},
	}
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
}

// SkuConfigDeleteOne is the builder for deleting a single SkuConfig entity.
type SkuConfigDeleteOne struct {
	scd *SkuConfigDelete
}

// Exec executes the deletion query.
func (scdo *SkuConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{skuconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SkuConfigDeleteOne) ExecX(ctx context.Context) {
	scdo.scd.ExecX(ctx)
}
