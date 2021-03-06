// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/predicate"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skuaction"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/skustatuscode"
)

// SkuStatusCodeQuery is the builder for querying SkuStatusCode entities.
type SkuStatusCodeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.SkuStatusCode
	// eager-loading edges.
	withAction *SkuActionQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (sscq *SkuStatusCodeQuery) Where(ps ...predicate.SkuStatusCode) *SkuStatusCodeQuery {
	sscq.predicates = append(sscq.predicates, ps...)
	return sscq
}

// Limit adds a limit step to the query.
func (sscq *SkuStatusCodeQuery) Limit(limit int) *SkuStatusCodeQuery {
	sscq.limit = &limit
	return sscq
}

// Offset adds an offset step to the query.
func (sscq *SkuStatusCodeQuery) Offset(offset int) *SkuStatusCodeQuery {
	sscq.offset = &offset
	return sscq
}

// Order adds an order step to the query.
func (sscq *SkuStatusCodeQuery) Order(o ...OrderFunc) *SkuStatusCodeQuery {
	sscq.order = append(sscq.order, o...)
	return sscq
}

// QueryAction chains the current query on the action edge.
func (sscq *SkuStatusCodeQuery) QueryAction() *SkuActionQuery {
	query := &SkuActionQuery{config: sscq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sscq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sscq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(skustatuscode.Table, skustatuscode.FieldID, selector),
			sqlgraph.To(skuaction.Table, skuaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, skustatuscode.ActionTable, skustatuscode.ActionColumn),
		)
		fromU = sqlgraph.SetNeighbors(sscq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SkuStatusCode entity in the query. Returns *NotFoundError when no skustatuscode was found.
func (sscq *SkuStatusCodeQuery) First(ctx context.Context) (*SkuStatusCode, error) {
	nodes, err := sscq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skustatuscode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) FirstX(ctx context.Context) *SkuStatusCode {
	node, err := sscq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SkuStatusCode id in the query. Returns *NotFoundError when no id was found.
func (sscq *SkuStatusCodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sscq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skustatuscode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) FirstIDX(ctx context.Context) int {
	id, err := sscq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only SkuStatusCode entity in the query, returns an error if not exactly one entity was returned.
func (sscq *SkuStatusCodeQuery) Only(ctx context.Context) (*SkuStatusCode, error) {
	nodes, err := sscq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skustatuscode.Label}
	default:
		return nil, &NotSingularError{skustatuscode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) OnlyX(ctx context.Context) *SkuStatusCode {
	node, err := sscq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only SkuStatusCode id in the query, returns an error if not exactly one id was returned.
func (sscq *SkuStatusCodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sscq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = &NotSingularError{skustatuscode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := sscq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SkuStatusCodes.
func (sscq *SkuStatusCodeQuery) All(ctx context.Context) ([]*SkuStatusCode, error) {
	if err := sscq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sscq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) AllX(ctx context.Context) []*SkuStatusCode {
	nodes, err := sscq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SkuStatusCode ids.
func (sscq *SkuStatusCodeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sscq.Select(skustatuscode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) IDsX(ctx context.Context) []int {
	ids, err := sscq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sscq *SkuStatusCodeQuery) Count(ctx context.Context) (int, error) {
	if err := sscq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sscq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) CountX(ctx context.Context) int {
	count, err := sscq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sscq *SkuStatusCodeQuery) Exist(ctx context.Context) (bool, error) {
	if err := sscq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sscq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sscq *SkuStatusCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := sscq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sscq *SkuStatusCodeQuery) Clone() *SkuStatusCodeQuery {
	if sscq == nil {
		return nil
	}
	return &SkuStatusCodeQuery{
		config:     sscq.config,
		limit:      sscq.limit,
		offset:     sscq.offset,
		order:      append([]OrderFunc{}, sscq.order...),
		predicates: append([]predicate.SkuStatusCode{}, sscq.predicates...),
		withAction: sscq.withAction.Clone(),
		// clone intermediate query.
		sql:  sscq.sql.Clone(),
		path: sscq.path,
	}
}

//  WithAction tells the query-builder to eager-loads the nodes that are connected to
// the "action" edge. The optional arguments used to configure the query builder of the edge.
func (sscq *SkuStatusCodeQuery) WithAction(opts ...func(*SkuActionQuery)) *SkuStatusCodeQuery {
	query := &SkuActionQuery{config: sscq.config}
	for _, opt := range opts {
		opt(query)
	}
	sscq.withAction = query
	return sscq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InsuranceCode string `json:"insurance_code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SkuStatusCode.Query().
//		GroupBy(skustatuscode.FieldInsuranceCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sscq *SkuStatusCodeQuery) GroupBy(field string, fields ...string) *SkuStatusCodeGroupBy {
	group := &SkuStatusCodeGroupBy{config: sscq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sscq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sscq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		InsuranceCode string `json:"insurance_code,omitempty"`
//	}
//
//	client.SkuStatusCode.Query().
//		Select(skustatuscode.FieldInsuranceCode).
//		Scan(ctx, &v)
//
func (sscq *SkuStatusCodeQuery) Select(field string, fields ...string) *SkuStatusCodeSelect {
	selector := &SkuStatusCodeSelect{config: sscq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sscq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sscq.sqlQuery(), nil
	}
	return selector
}

func (sscq *SkuStatusCodeQuery) prepareQuery(ctx context.Context) error {
	if sscq.path != nil {
		prev, err := sscq.path(ctx)
		if err != nil {
			return err
		}
		sscq.sql = prev
	}
	return nil
}

func (sscq *SkuStatusCodeQuery) sqlAll(ctx context.Context) ([]*SkuStatusCode, error) {
	var (
		nodes       = []*SkuStatusCode{}
		withFKs     = sscq.withFKs
		_spec       = sscq.querySpec()
		loadedTypes = [1]bool{
			sscq.withAction != nil,
		}
	)
	if sscq.withAction != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, skustatuscode.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &SkuStatusCode{config: sscq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, sscq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sscq.withAction; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*SkuStatusCode)
		for i := range nodes {
			if fk := nodes[i].sku_action_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(skuaction.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "sku_action_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Action = n
			}
		}
	}

	return nodes, nil
}

func (sscq *SkuStatusCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sscq.querySpec()
	return sqlgraph.CountNodes(ctx, sscq.driver, _spec)
}

func (sscq *SkuStatusCodeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sscq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sscq *SkuStatusCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skustatuscode.Table,
			Columns: skustatuscode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skustatuscode.FieldID,
			},
		},
		From:   sscq.sql,
		Unique: true,
	}
	if ps := sscq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sscq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sscq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sscq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, skustatuscode.ValidColumn)
			}
		}
	}
	return _spec
}

func (sscq *SkuStatusCodeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sscq.driver.Dialect())
	t1 := builder.Table(skustatuscode.Table)
	selector := builder.Select(t1.Columns(skustatuscode.Columns...)...).From(t1)
	if sscq.sql != nil {
		selector = sscq.sql
		selector.Select(selector.Columns(skustatuscode.Columns...)...)
	}
	for _, p := range sscq.predicates {
		p(selector)
	}
	for _, p := range sscq.order {
		p(selector, skustatuscode.ValidColumn)
	}
	if offset := sscq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sscq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SkuStatusCodeGroupBy is the builder for group-by SkuStatusCode entities.
type SkuStatusCodeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sscgb *SkuStatusCodeGroupBy) Aggregate(fns ...AggregateFunc) *SkuStatusCodeGroupBy {
	sscgb.fns = append(sscgb.fns, fns...)
	return sscgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sscgb *SkuStatusCodeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sscgb.path(ctx)
	if err != nil {
		return err
	}
	sscgb.sql = query
	return sscgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sscgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sscgb.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sscgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) StringsX(ctx context.Context) []string {
	v, err := sscgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sscgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) StringX(ctx context.Context) string {
	v, err := sscgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sscgb.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sscgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) IntsX(ctx context.Context) []int {
	v, err := sscgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sscgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) IntX(ctx context.Context) int {
	v, err := sscgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sscgb.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sscgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sscgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sscgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sscgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sscgb.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sscgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sscgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (sscgb *SkuStatusCodeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sscgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sscgb *SkuStatusCodeGroupBy) BoolX(ctx context.Context) bool {
	v, err := sscgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sscgb *SkuStatusCodeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sscgb.fields {
		if !skustatuscode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sscgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sscgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sscgb *SkuStatusCodeGroupBy) sqlQuery() *sql.Selector {
	selector := sscgb.sql
	columns := make([]string, 0, len(sscgb.fields)+len(sscgb.fns))
	columns = append(columns, sscgb.fields...)
	for _, fn := range sscgb.fns {
		columns = append(columns, fn(selector, skustatuscode.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(sscgb.fields...)
}

// SkuStatusCodeSelect is the builder for select fields of SkuStatusCode entities.
type SkuStatusCodeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (sscs *SkuStatusCodeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := sscs.path(ctx)
	if err != nil {
		return err
	}
	sscs.sql = query
	return sscs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sscs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sscs.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sscs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) StringsX(ctx context.Context) []string {
	v, err := sscs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sscs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) StringX(ctx context.Context) string {
	v, err := sscs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sscs.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sscs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) IntsX(ctx context.Context) []int {
	v, err := sscs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sscs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) IntX(ctx context.Context) int {
	v, err := sscs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sscs.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sscs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sscs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sscs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) Float64X(ctx context.Context) float64 {
	v, err := sscs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sscs.fields) > 1 {
		return nil, errors.New("ent: SkuStatusCodeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sscs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) BoolsX(ctx context.Context) []bool {
	v, err := sscs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (sscs *SkuStatusCodeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sscs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{skustatuscode.Label}
	default:
		err = fmt.Errorf("ent: SkuStatusCodeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sscs *SkuStatusCodeSelect) BoolX(ctx context.Context) bool {
	v, err := sscs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sscs *SkuStatusCodeSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sscs.fields {
		if !skustatuscode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := sscs.sqlQuery().Query()
	if err := sscs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sscs *SkuStatusCodeSelect) sqlQuery() sql.Querier {
	selector := sscs.sql
	selector.Select(selector.Columns(sscs.fields...)...)
	return selector
}
