// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningprofitdetail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// MiningProfitDetailQuery is the builder for querying MiningProfitDetail entities.
type MiningProfitDetailQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MiningProfitDetail
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MiningProfitDetailQuery builder.
func (mpdq *MiningProfitDetailQuery) Where(ps ...predicate.MiningProfitDetail) *MiningProfitDetailQuery {
	mpdq.predicates = append(mpdq.predicates, ps...)
	return mpdq
}

// Limit adds a limit step to the query.
func (mpdq *MiningProfitDetailQuery) Limit(limit int) *MiningProfitDetailQuery {
	mpdq.limit = &limit
	return mpdq
}

// Offset adds an offset step to the query.
func (mpdq *MiningProfitDetailQuery) Offset(offset int) *MiningProfitDetailQuery {
	mpdq.offset = &offset
	return mpdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpdq *MiningProfitDetailQuery) Unique(unique bool) *MiningProfitDetailQuery {
	mpdq.unique = &unique
	return mpdq
}

// Order adds an order step to the query.
func (mpdq *MiningProfitDetailQuery) Order(o ...OrderFunc) *MiningProfitDetailQuery {
	mpdq.order = append(mpdq.order, o...)
	return mpdq
}

// First returns the first MiningProfitDetail entity from the query.
// Returns a *NotFoundError when no MiningProfitDetail was found.
func (mpdq *MiningProfitDetailQuery) First(ctx context.Context) (*MiningProfitDetail, error) {
	nodes, err := mpdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{miningprofitdetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) FirstX(ctx context.Context) *MiningProfitDetail {
	node, err := mpdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MiningProfitDetail ID from the query.
// Returns a *NotFoundError when no MiningProfitDetail ID was found.
func (mpdq *MiningProfitDetailQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mpdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{miningprofitdetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mpdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MiningProfitDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MiningProfitDetail entity is found.
// Returns a *NotFoundError when no MiningProfitDetail entities are found.
func (mpdq *MiningProfitDetailQuery) Only(ctx context.Context) (*MiningProfitDetail, error) {
	nodes, err := mpdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{miningprofitdetail.Label}
	default:
		return nil, &NotSingularError{miningprofitdetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) OnlyX(ctx context.Context) *MiningProfitDetail {
	node, err := mpdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MiningProfitDetail ID in the query.
// Returns a *NotSingularError when more than one MiningProfitDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpdq *MiningProfitDetailQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mpdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{miningprofitdetail.Label}
	default:
		err = &NotSingularError{miningprofitdetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mpdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MiningProfitDetails.
func (mpdq *MiningProfitDetailQuery) All(ctx context.Context) ([]*MiningProfitDetail, error) {
	if err := mpdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mpdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) AllX(ctx context.Context) []*MiningProfitDetail {
	nodes, err := mpdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MiningProfitDetail IDs.
func (mpdq *MiningProfitDetailQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := mpdq.Select(miningprofitdetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mpdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpdq *MiningProfitDetailQuery) Count(ctx context.Context) (int, error) {
	if err := mpdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mpdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) CountX(ctx context.Context) int {
	count, err := mpdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpdq *MiningProfitDetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := mpdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mpdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mpdq *MiningProfitDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := mpdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MiningProfitDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpdq *MiningProfitDetailQuery) Clone() *MiningProfitDetailQuery {
	if mpdq == nil {
		return nil
	}
	return &MiningProfitDetailQuery{
		config:     mpdq.config,
		limit:      mpdq.limit,
		offset:     mpdq.offset,
		order:      append([]OrderFunc{}, mpdq.order...),
		predicates: append([]predicate.MiningProfitDetail{}, mpdq.predicates...),
		// clone intermediate query.
		sql:    mpdq.sql.Clone(),
		path:   mpdq.path,
		unique: mpdq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MiningProfitDetail.Query().
//		GroupBy(miningprofitdetail.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mpdq *MiningProfitDetailQuery) GroupBy(field string, fields ...string) *MiningProfitDetailGroupBy {
	grbuild := &MiningProfitDetailGroupBy{config: mpdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mpdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mpdq.sqlQuery(ctx), nil
	}
	grbuild.label = miningprofitdetail.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.MiningProfitDetail.Query().
//		Select(miningprofitdetail.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (mpdq *MiningProfitDetailQuery) Select(fields ...string) *MiningProfitDetailSelect {
	mpdq.fields = append(mpdq.fields, fields...)
	selbuild := &MiningProfitDetailSelect{MiningProfitDetailQuery: mpdq}
	selbuild.label = miningprofitdetail.Label
	selbuild.flds, selbuild.scan = &mpdq.fields, selbuild.Scan
	return selbuild
}

func (mpdq *MiningProfitDetailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mpdq.fields {
		if !miningprofitdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mpdq.path != nil {
		prev, err := mpdq.path(ctx)
		if err != nil {
			return err
		}
		mpdq.sql = prev
	}
	if miningprofitdetail.Policy == nil {
		return errors.New("ent: uninitialized miningprofitdetail.Policy (forgotten import ent/runtime?)")
	}
	if err := miningprofitdetail.Policy.EvalQuery(ctx, mpdq); err != nil {
		return err
	}
	return nil
}

func (mpdq *MiningProfitDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MiningProfitDetail, error) {
	var (
		nodes = []*MiningProfitDetail{}
		_spec = mpdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*MiningProfitDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &MiningProfitDetail{config: mpdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(mpdq.modifiers) > 0 {
		_spec.Modifiers = mpdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mpdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mpdq *MiningProfitDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpdq.querySpec()
	if len(mpdq.modifiers) > 0 {
		_spec.Modifiers = mpdq.modifiers
	}
	_spec.Node.Columns = mpdq.fields
	if len(mpdq.fields) > 0 {
		_spec.Unique = mpdq.unique != nil && *mpdq.unique
	}
	return sqlgraph.CountNodes(ctx, mpdq.driver, _spec)
}

func (mpdq *MiningProfitDetailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mpdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mpdq *MiningProfitDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   miningprofitdetail.Table,
			Columns: miningprofitdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: miningprofitdetail.FieldID,
			},
		},
		From:   mpdq.sql,
		Unique: true,
	}
	if unique := mpdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mpdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, miningprofitdetail.FieldID)
		for i := range fields {
			if fields[i] != miningprofitdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mpdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mpdq *MiningProfitDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpdq.driver.Dialect())
	t1 := builder.Table(miningprofitdetail.Table)
	columns := mpdq.fields
	if len(columns) == 0 {
		columns = miningprofitdetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpdq.sql != nil {
		selector = mpdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpdq.unique != nil && *mpdq.unique {
		selector.Distinct()
	}
	for _, m := range mpdq.modifiers {
		m(selector)
	}
	for _, p := range mpdq.predicates {
		p(selector)
	}
	for _, p := range mpdq.order {
		p(selector)
	}
	if offset := mpdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (mpdq *MiningProfitDetailQuery) ForUpdate(opts ...sql.LockOption) *MiningProfitDetailQuery {
	if mpdq.driver.Dialect() == dialect.Postgres {
		mpdq.Unique(false)
	}
	mpdq.modifiers = append(mpdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return mpdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (mpdq *MiningProfitDetailQuery) ForShare(opts ...sql.LockOption) *MiningProfitDetailQuery {
	if mpdq.driver.Dialect() == dialect.Postgres {
		mpdq.Unique(false)
	}
	mpdq.modifiers = append(mpdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return mpdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mpdq *MiningProfitDetailQuery) Modify(modifiers ...func(s *sql.Selector)) *MiningProfitDetailSelect {
	mpdq.modifiers = append(mpdq.modifiers, modifiers...)
	return mpdq.Select()
}

// MiningProfitDetailGroupBy is the group-by builder for MiningProfitDetail entities.
type MiningProfitDetailGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpdgb *MiningProfitDetailGroupBy) Aggregate(fns ...AggregateFunc) *MiningProfitDetailGroupBy {
	mpdgb.fns = append(mpdgb.fns, fns...)
	return mpdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mpdgb *MiningProfitDetailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mpdgb.path(ctx)
	if err != nil {
		return err
	}
	mpdgb.sql = query
	return mpdgb.sqlScan(ctx, v)
}

func (mpdgb *MiningProfitDetailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mpdgb.fields {
		if !miningprofitdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mpdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mpdgb *MiningProfitDetailGroupBy) sqlQuery() *sql.Selector {
	selector := mpdgb.sql.Select()
	aggregation := make([]string, 0, len(mpdgb.fns))
	for _, fn := range mpdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mpdgb.fields)+len(mpdgb.fns))
		for _, f := range mpdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mpdgb.fields...)...)
}

// MiningProfitDetailSelect is the builder for selecting fields of MiningProfitDetail entities.
type MiningProfitDetailSelect struct {
	*MiningProfitDetailQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mpds *MiningProfitDetailSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mpds.prepareQuery(ctx); err != nil {
		return err
	}
	mpds.sql = mpds.MiningProfitDetailQuery.sqlQuery(ctx)
	return mpds.sqlScan(ctx, v)
}

func (mpds *MiningProfitDetailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mpds.sql.Query()
	if err := mpds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mpds *MiningProfitDetailSelect) Modify(modifiers ...func(s *sql.Selector)) *MiningProfitDetailSelect {
	mpds.modifiers = append(mpds.modifiers, modifiers...)
	return mpds
}
