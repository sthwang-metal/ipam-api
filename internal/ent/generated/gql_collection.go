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

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
	"go.infratographer.com/x/gidx"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ia *IPAddressQuery) CollectFields(ctx context.Context, satisfies ...string) (*IPAddressQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ia, nil
	}
	if err := ia.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ia, nil
}

func (ia *IPAddressQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(ipaddress.Columns))
		selectedFields = []string{ipaddress.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "ipBlock":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&IPBlockClient{config: ia.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ia.withIPBlock = query
			if _, ok := fieldSeen[ipaddress.FieldBlockID]; !ok {
				selectedFields = append(selectedFields, ipaddress.FieldBlockID)
				fieldSeen[ipaddress.FieldBlockID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[ipaddress.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, ipaddress.FieldCreatedAt)
				fieldSeen[ipaddress.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[ipaddress.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, ipaddress.FieldUpdatedAt)
				fieldSeen[ipaddress.FieldUpdatedAt] = struct{}{}
			}
		case "ip":
			if _, ok := fieldSeen[ipaddress.FieldIP]; !ok {
				selectedFields = append(selectedFields, ipaddress.FieldIP)
				fieldSeen[ipaddress.FieldIP] = struct{}{}
			}
		case "reserved":
			if _, ok := fieldSeen[ipaddress.FieldReserved]; !ok {
				selectedFields = append(selectedFields, ipaddress.FieldReserved)
				fieldSeen[ipaddress.FieldReserved] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ia.Select(selectedFields...)
	}
	return nil
}

type ipaddressPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []IPAddressPaginateOption
}

func newIPAddressPaginateArgs(rv map[string]any) *ipaddressPaginateArgs {
	args := &ipaddressPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &IPAddressOrder{Field: &IPAddressOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithIPAddressOrder(order))
			}
		case *IPAddressOrder:
			if v != nil {
				args.opts = append(args.opts, WithIPAddressOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*IPAddressWhereInput); ok {
		args.opts = append(args.opts, WithIPAddressFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ib *IPBlockQuery) CollectFields(ctx context.Context, satisfies ...string) (*IPBlockQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ib, nil
	}
	if err := ib.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ib, nil
}

func (ib *IPBlockQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(ipblock.Columns))
		selectedFields = []string{ipblock.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "ipBlockType":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&IPBlockTypeClient{config: ib.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ib.withIPBlockType = query
			if _, ok := fieldSeen[ipblock.FieldBlockTypeID]; !ok {
				selectedFields = append(selectedFields, ipblock.FieldBlockTypeID)
				fieldSeen[ipblock.FieldBlockTypeID] = struct{}{}
			}
		case "ipAddress":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&IPAddressClient{config: ib.config}).Query()
			)
			args := newIPAddressPaginateArgs(fieldArgs(ctx, new(IPAddressWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newIPAddressPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					ib.loadTotal = append(ib.loadTotal, func(ctx context.Context, nodes []*IPBlock) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID gidx.PrefixedID `sql:"block_id"`
							Count  int             `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(ipblock.IPAddressColumn), ids...))
						})
						if err := query.GroupBy(ipblock.IPAddressColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[gidx.PrefixedID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					ib.loadTotal = append(ib.loadTotal, func(_ context.Context, nodes []*IPBlock) error {
						for i := range nodes {
							n := len(nodes[i].Edges.IPAddress)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "IPAddress")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(ipblock.IPAddressColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			ib.WithNamedIPAddress(alias, func(wq *IPAddressQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[ipblock.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, ipblock.FieldCreatedAt)
				fieldSeen[ipblock.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[ipblock.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, ipblock.FieldUpdatedAt)
				fieldSeen[ipblock.FieldUpdatedAt] = struct{}{}
			}
		case "prefix":
			if _, ok := fieldSeen[ipblock.FieldPrefix]; !ok {
				selectedFields = append(selectedFields, ipblock.FieldPrefix)
				fieldSeen[ipblock.FieldPrefix] = struct{}{}
			}
		case "allowAutoSubnet":
			if _, ok := fieldSeen[ipblock.FieldAllowAutoSubnet]; !ok {
				selectedFields = append(selectedFields, ipblock.FieldAllowAutoSubnet)
				fieldSeen[ipblock.FieldAllowAutoSubnet] = struct{}{}
			}
		case "allowAutoAllocate":
			if _, ok := fieldSeen[ipblock.FieldAllowAutoAllocate]; !ok {
				selectedFields = append(selectedFields, ipblock.FieldAllowAutoAllocate)
				fieldSeen[ipblock.FieldAllowAutoAllocate] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ib.Select(selectedFields...)
	}
	return nil
}

type ipblockPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []IPBlockPaginateOption
}

func newIPBlockPaginateArgs(rv map[string]any) *ipblockPaginateArgs {
	args := &ipblockPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &IPBlockOrder{Field: &IPBlockOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithIPBlockOrder(order))
			}
		case *IPBlockOrder:
			if v != nil {
				args.opts = append(args.opts, WithIPBlockOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*IPBlockWhereInput); ok {
		args.opts = append(args.opts, WithIPBlockFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ibt *IPBlockTypeQuery) CollectFields(ctx context.Context, satisfies ...string) (*IPBlockTypeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ibt, nil
	}
	if err := ibt.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ibt, nil
}

func (ibt *IPBlockTypeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(ipblocktype.Columns))
		selectedFields = []string{ipblocktype.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "ipBlock":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&IPBlockClient{config: ibt.config}).Query()
			)
			args := newIPBlockPaginateArgs(fieldArgs(ctx, new(IPBlockWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newIPBlockPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					ibt.loadTotal = append(ibt.loadTotal, func(ctx context.Context, nodes []*IPBlockType) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID gidx.PrefixedID `sql:"block_type_id"`
							Count  int             `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(ipblocktype.IPBlockColumn), ids...))
						})
						if err := query.GroupBy(ipblocktype.IPBlockColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[gidx.PrefixedID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					ibt.loadTotal = append(ibt.loadTotal, func(_ context.Context, nodes []*IPBlockType) error {
						for i := range nodes {
							n := len(nodes[i].Edges.IPBlock)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "IPBlock")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(ipblocktype.IPBlockColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			ibt.WithNamedIPBlock(alias, func(wq *IPBlockQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[ipblocktype.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, ipblocktype.FieldCreatedAt)
				fieldSeen[ipblocktype.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[ipblocktype.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, ipblocktype.FieldUpdatedAt)
				fieldSeen[ipblocktype.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[ipblocktype.FieldName]; !ok {
				selectedFields = append(selectedFields, ipblocktype.FieldName)
				fieldSeen[ipblocktype.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ibt.Select(selectedFields...)
	}
	return nil
}

type ipblocktypePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []IPBlockTypePaginateOption
}

func newIPBlockTypePaginateArgs(rv map[string]any) *ipblocktypePaginateArgs {
	args := &ipblocktypePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &IPBlockTypeOrder{Field: &IPBlockTypeOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithIPBlockTypeOrder(order))
			}
		case *IPBlockTypeOrder:
			if v != nil {
				args.opts = append(args.opts, WithIPBlockTypeOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*IPBlockTypeWhereInput); ok {
		args.opts = append(args.opts, WithIPBlockTypeFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
