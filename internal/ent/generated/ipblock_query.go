// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
	"go.infratographer.com/ipam-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// IPBlockQuery is the builder for querying IPBlock entities.
type IPBlockQuery struct {
	config
	ctx                *QueryContext
	order              []ipblock.OrderOption
	inters             []Interceptor
	predicates         []predicate.IPBlock
	withIPBlockType    *IPBlockTypeQuery
	withIPAddress      *IPAddressQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*IPBlock) error
	withNamedIPAddress map[string]*IPAddressQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IPBlockQuery builder.
func (ibq *IPBlockQuery) Where(ps ...predicate.IPBlock) *IPBlockQuery {
	ibq.predicates = append(ibq.predicates, ps...)
	return ibq
}

// Limit the number of records to be returned by this query.
func (ibq *IPBlockQuery) Limit(limit int) *IPBlockQuery {
	ibq.ctx.Limit = &limit
	return ibq
}

// Offset to start from.
func (ibq *IPBlockQuery) Offset(offset int) *IPBlockQuery {
	ibq.ctx.Offset = &offset
	return ibq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ibq *IPBlockQuery) Unique(unique bool) *IPBlockQuery {
	ibq.ctx.Unique = &unique
	return ibq
}

// Order specifies how the records should be ordered.
func (ibq *IPBlockQuery) Order(o ...ipblock.OrderOption) *IPBlockQuery {
	ibq.order = append(ibq.order, o...)
	return ibq
}

// QueryIPBlockType chains the current query on the "ip_block_type" edge.
func (ibq *IPBlockQuery) QueryIPBlockType() *IPBlockTypeQuery {
	query := (&IPBlockTypeClient{config: ibq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ibq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ibq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ipblock.Table, ipblock.FieldID, selector),
			sqlgraph.To(ipblocktype.Table, ipblocktype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ipblock.IPBlockTypeTable, ipblock.IPBlockTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ibq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIPAddress chains the current query on the "ip_address" edge.
func (ibq *IPBlockQuery) QueryIPAddress() *IPAddressQuery {
	query := (&IPAddressClient{config: ibq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ibq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ibq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ipblock.Table, ipblock.FieldID, selector),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ipblock.IPAddressTable, ipblock.IPAddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(ibq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IPBlock entity from the query.
// Returns a *NotFoundError when no IPBlock was found.
func (ibq *IPBlockQuery) First(ctx context.Context) (*IPBlock, error) {
	nodes, err := ibq.Limit(1).All(setContextOp(ctx, ibq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ipblock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ibq *IPBlockQuery) FirstX(ctx context.Context) *IPBlock {
	node, err := ibq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IPBlock ID from the query.
// Returns a *NotFoundError when no IPBlock ID was found.
func (ibq *IPBlockQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = ibq.Limit(1).IDs(setContextOp(ctx, ibq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ipblock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ibq *IPBlockQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := ibq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IPBlock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IPBlock entity is found.
// Returns a *NotFoundError when no IPBlock entities are found.
func (ibq *IPBlockQuery) Only(ctx context.Context) (*IPBlock, error) {
	nodes, err := ibq.Limit(2).All(setContextOp(ctx, ibq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ipblock.Label}
	default:
		return nil, &NotSingularError{ipblock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ibq *IPBlockQuery) OnlyX(ctx context.Context) *IPBlock {
	node, err := ibq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IPBlock ID in the query.
// Returns a *NotSingularError when more than one IPBlock ID is found.
// Returns a *NotFoundError when no entities are found.
func (ibq *IPBlockQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = ibq.Limit(2).IDs(setContextOp(ctx, ibq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ipblock.Label}
	default:
		err = &NotSingularError{ipblock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ibq *IPBlockQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := ibq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IPBlocks.
func (ibq *IPBlockQuery) All(ctx context.Context) ([]*IPBlock, error) {
	ctx = setContextOp(ctx, ibq.ctx, "All")
	if err := ibq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IPBlock, *IPBlockQuery]()
	return withInterceptors[[]*IPBlock](ctx, ibq, qr, ibq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ibq *IPBlockQuery) AllX(ctx context.Context) []*IPBlock {
	nodes, err := ibq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IPBlock IDs.
func (ibq *IPBlockQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if ibq.ctx.Unique == nil && ibq.path != nil {
		ibq.Unique(true)
	}
	ctx = setContextOp(ctx, ibq.ctx, "IDs")
	if err = ibq.Select(ipblock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ibq *IPBlockQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := ibq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ibq *IPBlockQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ibq.ctx, "Count")
	if err := ibq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ibq, querierCount[*IPBlockQuery](), ibq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ibq *IPBlockQuery) CountX(ctx context.Context) int {
	count, err := ibq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ibq *IPBlockQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ibq.ctx, "Exist")
	switch _, err := ibq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ibq *IPBlockQuery) ExistX(ctx context.Context) bool {
	exist, err := ibq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IPBlockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ibq *IPBlockQuery) Clone() *IPBlockQuery {
	if ibq == nil {
		return nil
	}
	return &IPBlockQuery{
		config:          ibq.config,
		ctx:             ibq.ctx.Clone(),
		order:           append([]ipblock.OrderOption{}, ibq.order...),
		inters:          append([]Interceptor{}, ibq.inters...),
		predicates:      append([]predicate.IPBlock{}, ibq.predicates...),
		withIPBlockType: ibq.withIPBlockType.Clone(),
		withIPAddress:   ibq.withIPAddress.Clone(),
		// clone intermediate query.
		sql:  ibq.sql.Clone(),
		path: ibq.path,
	}
}

// WithIPBlockType tells the query-builder to eager-load the nodes that are connected to
// the "ip_block_type" edge. The optional arguments are used to configure the query builder of the edge.
func (ibq *IPBlockQuery) WithIPBlockType(opts ...func(*IPBlockTypeQuery)) *IPBlockQuery {
	query := (&IPBlockTypeClient{config: ibq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ibq.withIPBlockType = query
	return ibq
}

// WithIPAddress tells the query-builder to eager-load the nodes that are connected to
// the "ip_address" edge. The optional arguments are used to configure the query builder of the edge.
func (ibq *IPBlockQuery) WithIPAddress(opts ...func(*IPAddressQuery)) *IPBlockQuery {
	query := (&IPAddressClient{config: ibq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ibq.withIPAddress = query
	return ibq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IPBlock.Query().
//		GroupBy(ipblock.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ibq *IPBlockQuery) GroupBy(field string, fields ...string) *IPBlockGroupBy {
	ibq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IPBlockGroupBy{build: ibq}
	grbuild.flds = &ibq.ctx.Fields
	grbuild.label = ipblock.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.IPBlock.Query().
//		Select(ipblock.FieldCreatedAt).
//		Scan(ctx, &v)
func (ibq *IPBlockQuery) Select(fields ...string) *IPBlockSelect {
	ibq.ctx.Fields = append(ibq.ctx.Fields, fields...)
	sbuild := &IPBlockSelect{IPBlockQuery: ibq}
	sbuild.label = ipblock.Label
	sbuild.flds, sbuild.scan = &ibq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IPBlockSelect configured with the given aggregations.
func (ibq *IPBlockQuery) Aggregate(fns ...AggregateFunc) *IPBlockSelect {
	return ibq.Select().Aggregate(fns...)
}

func (ibq *IPBlockQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ibq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ibq); err != nil {
				return err
			}
		}
	}
	for _, f := range ibq.ctx.Fields {
		if !ipblock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ibq.path != nil {
		prev, err := ibq.path(ctx)
		if err != nil {
			return err
		}
		ibq.sql = prev
	}
	return nil
}

func (ibq *IPBlockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IPBlock, error) {
	var (
		nodes       = []*IPBlock{}
		_spec       = ibq.querySpec()
		loadedTypes = [2]bool{
			ibq.withIPBlockType != nil,
			ibq.withIPAddress != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IPBlock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IPBlock{config: ibq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ibq.modifiers) > 0 {
		_spec.Modifiers = ibq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ibq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ibq.withIPBlockType; query != nil {
		if err := ibq.loadIPBlockType(ctx, query, nodes, nil,
			func(n *IPBlock, e *IPBlockType) { n.Edges.IPBlockType = e }); err != nil {
			return nil, err
		}
	}
	if query := ibq.withIPAddress; query != nil {
		if err := ibq.loadIPAddress(ctx, query, nodes,
			func(n *IPBlock) { n.Edges.IPAddress = []*IPAddress{} },
			func(n *IPBlock, e *IPAddress) { n.Edges.IPAddress = append(n.Edges.IPAddress, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ibq.withNamedIPAddress {
		if err := ibq.loadIPAddress(ctx, query, nodes,
			func(n *IPBlock) { n.appendNamedIPAddress(name) },
			func(n *IPBlock, e *IPAddress) { n.appendNamedIPAddress(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ibq.loadTotal {
		if err := ibq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ibq *IPBlockQuery) loadIPBlockType(ctx context.Context, query *IPBlockTypeQuery, nodes []*IPBlock, init func(*IPBlock), assign func(*IPBlock, *IPBlockType)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*IPBlock)
	for i := range nodes {
		fk := nodes[i].BlockTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ipblocktype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "block_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ibq *IPBlockQuery) loadIPAddress(ctx context.Context, query *IPAddressQuery, nodes []*IPBlock, init func(*IPBlock), assign func(*IPBlock, *IPAddress)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID]*IPBlock)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(ipaddress.FieldBlockID)
	}
	query.Where(predicate.IPAddress(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(ipblock.IPAddressColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BlockID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "block_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ibq *IPBlockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ibq.querySpec()
	if len(ibq.modifiers) > 0 {
		_spec.Modifiers = ibq.modifiers
	}
	_spec.Node.Columns = ibq.ctx.Fields
	if len(ibq.ctx.Fields) > 0 {
		_spec.Unique = ibq.ctx.Unique != nil && *ibq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ibq.driver, _spec)
}

func (ibq *IPBlockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ipblock.Table, ipblock.Columns, sqlgraph.NewFieldSpec(ipblock.FieldID, field.TypeString))
	_spec.From = ibq.sql
	if unique := ibq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ibq.path != nil {
		_spec.Unique = true
	}
	if fields := ibq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipblock.FieldID)
		for i := range fields {
			if fields[i] != ipblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ibq.withIPBlockType != nil {
			_spec.Node.AddColumnOnce(ipblock.FieldBlockTypeID)
		}
	}
	if ps := ibq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ibq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ibq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ibq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ibq *IPBlockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ibq.driver.Dialect())
	t1 := builder.Table(ipblock.Table)
	columns := ibq.ctx.Fields
	if len(columns) == 0 {
		columns = ipblock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ibq.sql != nil {
		selector = ibq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ibq.ctx.Unique != nil && *ibq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ibq.predicates {
		p(selector)
	}
	for _, p := range ibq.order {
		p(selector)
	}
	if offset := ibq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ibq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedIPAddress tells the query-builder to eager-load the nodes that are connected to the "ip_address"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ibq *IPBlockQuery) WithNamedIPAddress(name string, opts ...func(*IPAddressQuery)) *IPBlockQuery {
	query := (&IPAddressClient{config: ibq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ibq.withNamedIPAddress == nil {
		ibq.withNamedIPAddress = make(map[string]*IPAddressQuery)
	}
	ibq.withNamedIPAddress[name] = query
	return ibq
}

// IPBlockGroupBy is the group-by builder for IPBlock entities.
type IPBlockGroupBy struct {
	selector
	build *IPBlockQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ibgb *IPBlockGroupBy) Aggregate(fns ...AggregateFunc) *IPBlockGroupBy {
	ibgb.fns = append(ibgb.fns, fns...)
	return ibgb
}

// Scan applies the selector query and scans the result into the given value.
func (ibgb *IPBlockGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ibgb.build.ctx, "GroupBy")
	if err := ibgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IPBlockQuery, *IPBlockGroupBy](ctx, ibgb.build, ibgb, ibgb.build.inters, v)
}

func (ibgb *IPBlockGroupBy) sqlScan(ctx context.Context, root *IPBlockQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ibgb.fns))
	for _, fn := range ibgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ibgb.flds)+len(ibgb.fns))
		for _, f := range *ibgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ibgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ibgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IPBlockSelect is the builder for selecting fields of IPBlock entities.
type IPBlockSelect struct {
	*IPBlockQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ibs *IPBlockSelect) Aggregate(fns ...AggregateFunc) *IPBlockSelect {
	ibs.fns = append(ibs.fns, fns...)
	return ibs
}

// Scan applies the selector query and scans the result into the given value.
func (ibs *IPBlockSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ibs.ctx, "Select")
	if err := ibs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IPBlockQuery, *IPBlockSelect](ctx, ibs.IPBlockQuery, ibs, ibs.inters, v)
}

func (ibs *IPBlockSelect) sqlScan(ctx context.Context, root *IPBlockQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ibs.fns))
	for _, fn := range ibs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ibs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ibs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
