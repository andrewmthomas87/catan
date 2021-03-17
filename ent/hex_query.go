// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/hex"
	"github.com/andrewmthomas87/catan/ent/numbertoken"
	"github.com/andrewmthomas87/catan/ent/predicate"
	"github.com/andrewmthomas87/catan/ent/robber"
	"github.com/andrewmthomas87/catan/ent/settlement"
)

// HexQuery is the builder for querying Hex entities.
type HexQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Hex
	// eager-loading edges.
	withNumberToken *NumberTokenQuery
	withRobber      *RobberQuery
	withSettlements *SettlementQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HexQuery builder.
func (hq *HexQuery) Where(ps ...predicate.Hex) *HexQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HexQuery) Limit(limit int) *HexQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HexQuery) Offset(offset int) *HexQuery {
	hq.offset = &offset
	return hq
}

// Order adds an order step to the query.
func (hq *HexQuery) Order(o ...OrderFunc) *HexQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryNumberToken chains the current query on the "number_token" edge.
func (hq *HexQuery) QueryNumberToken() *NumberTokenQuery {
	query := &NumberTokenQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hex.Table, hex.FieldID, selector),
			sqlgraph.To(numbertoken.Table, numbertoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, hex.NumberTokenTable, hex.NumberTokenColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRobber chains the current query on the "robber" edge.
func (hq *HexQuery) QueryRobber() *RobberQuery {
	query := &RobberQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hex.Table, hex.FieldID, selector),
			sqlgraph.To(robber.Table, robber.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, hex.RobberTable, hex.RobberColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySettlements chains the current query on the "settlements" edge.
func (hq *HexQuery) QuerySettlements() *SettlementQuery {
	query := &SettlementQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hex.Table, hex.FieldID, selector),
			sqlgraph.To(settlement.Table, settlement.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, hex.SettlementsTable, hex.SettlementsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Hex entity from the query.
// Returns a *NotFoundError when no Hex was found.
func (hq *HexQuery) First(ctx context.Context) (*Hex, error) {
	nodes, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hex.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HexQuery) FirstX(ctx context.Context) *Hex {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Hex ID from the query.
// Returns a *NotFoundError when no Hex ID was found.
func (hq *HexQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hex.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HexQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Hex entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Hex entity is not found.
// Returns a *NotFoundError when no Hex entities are found.
func (hq *HexQuery) Only(ctx context.Context) (*Hex, error) {
	nodes, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hex.Label}
	default:
		return nil, &NotSingularError{hex.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HexQuery) OnlyX(ctx context.Context) *Hex {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Hex ID in the query.
// Returns a *NotSingularError when exactly one Hex ID is not found.
// Returns a *NotFoundError when no entities are found.
func (hq *HexQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = &NotSingularError{hex.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HexQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Hexes.
func (hq *HexQuery) All(ctx context.Context) ([]*Hex, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HexQuery) AllX(ctx context.Context) []*Hex {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Hex IDs.
func (hq *HexQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hq.Select(hex.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HexQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HexQuery) Count(ctx context.Context) (int, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HexQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HexQuery) Exist(ctx context.Context) (bool, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HexQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HexQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HexQuery) Clone() *HexQuery {
	if hq == nil {
		return nil
	}
	return &HexQuery{
		config:          hq.config,
		limit:           hq.limit,
		offset:          hq.offset,
		order:           append([]OrderFunc{}, hq.order...),
		predicates:      append([]predicate.Hex{}, hq.predicates...),
		withNumberToken: hq.withNumberToken.Clone(),
		withRobber:      hq.withRobber.Clone(),
		withSettlements: hq.withSettlements.Clone(),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// WithNumberToken tells the query-builder to eager-load the nodes that are connected to
// the "number_token" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HexQuery) WithNumberToken(opts ...func(*NumberTokenQuery)) *HexQuery {
	query := &NumberTokenQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withNumberToken = query
	return hq
}

// WithRobber tells the query-builder to eager-load the nodes that are connected to
// the "robber" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HexQuery) WithRobber(opts ...func(*RobberQuery)) *HexQuery {
	query := &RobberQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withRobber = query
	return hq
}

// WithSettlements tells the query-builder to eager-load the nodes that are connected to
// the "settlements" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HexQuery) WithSettlements(opts ...func(*SettlementQuery)) *HexQuery {
	query := &SettlementQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withSettlements = query
	return hq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		X int `json:"x,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Hex.Query().
//		GroupBy(hex.FieldX).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hq *HexQuery) GroupBy(field string, fields ...string) *HexGroupBy {
	group := &HexGroupBy{config: hq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		X int `json:"x,omitempty"`
//	}
//
//	client.Hex.Query().
//		Select(hex.FieldX).
//		Scan(ctx, &v)
//
func (hq *HexQuery) Select(field string, fields ...string) *HexSelect {
	hq.fields = append([]string{field}, fields...)
	return &HexSelect{HexQuery: hq}
}

func (hq *HexQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hq.fields {
		if !hex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HexQuery) sqlAll(ctx context.Context) ([]*Hex, error) {
	var (
		nodes       = []*Hex{}
		withFKs     = hq.withFKs
		_spec       = hq.querySpec()
		loadedTypes = [3]bool{
			hq.withNumberToken != nil,
			hq.withRobber != nil,
			hq.withSettlements != nil,
		}
	)
	if hq.withNumberToken != nil || hq.withRobber != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, hex.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Hex{config: hq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := hq.withNumberToken; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Hex)
		for i := range nodes {
			if fk := nodes[i].number_token_hex; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(numbertoken.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "number_token_hex" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.NumberToken = n
			}
		}
	}

	if query := hq.withRobber; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Hex)
		for i := range nodes {
			if fk := nodes[i].robber_hex; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(robber.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "robber_hex" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Robber = n
			}
		}
	}

	if query := hq.withSettlements; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Hex, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Settlements = []*Settlement{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Hex)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   hex.SettlementsTable,
				Columns: hex.SettlementsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(hex.SettlementsPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, hq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "settlements": %v`, err)
		}
		query.Where(settlement.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "settlements" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Settlements = append(nodes[i].Edges.Settlements, n)
			}
		}
	}

	return nodes, nil
}

func (hq *HexQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HexQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (hq *HexQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hex.Table,
			Columns: hex.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hex.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if fields := hq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hex.FieldID)
		for i := range fields {
			if fields[i] != hex.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, hex.ValidColumn)
			}
		}
	}
	return _spec
}

func (hq *HexQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(hex.Table)
	selector := builder.Select(t1.Columns(hex.Columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(hex.Columns...)...)
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector, hex.ValidColumn)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HexGroupBy is the group-by builder for Hex entities.
type HexGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HexGroupBy) Aggregate(fns ...AggregateFunc) *HexGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hgb *HexGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hgb.path(ctx)
	if err != nil {
		return err
	}
	hgb.sql = query
	return hgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hgb *HexGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HexGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hgb *HexGroupBy) StringsX(ctx context.Context) []string {
	v, err := hgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hgb *HexGroupBy) StringX(ctx context.Context) string {
	v, err := hgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HexGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hgb *HexGroupBy) IntsX(ctx context.Context) []int {
	v, err := hgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hgb *HexGroupBy) IntX(ctx context.Context) int {
	v, err := hgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HexGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hgb *HexGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hgb *HexGroupBy) Float64X(ctx context.Context) float64 {
	v, err := hgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HexGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hgb *HexGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hgb *HexGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hgb *HexGroupBy) BoolX(ctx context.Context) bool {
	v, err := hgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hgb *HexGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hgb.fields {
		if !hex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HexGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql
	columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
	columns = append(columns, hgb.fields...)
	for _, fn := range hgb.fns {
		columns = append(columns, fn(selector, hex.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(hgb.fields...)
}

// HexSelect is the builder for selecting fields of Hex entities.
type HexSelect struct {
	*HexQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HexSelect) Scan(ctx context.Context, v interface{}) error {
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	hs.sql = hs.HexQuery.sqlQuery(ctx)
	return hs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hs *HexSelect) ScanX(ctx context.Context, v interface{}) {
	if err := hs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Strings(ctx context.Context) ([]string, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HexSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hs *HexSelect) StringsX(ctx context.Context) []string {
	v, err := hs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hs *HexSelect) StringX(ctx context.Context) string {
	v, err := hs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Ints(ctx context.Context) ([]int, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HexSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hs *HexSelect) IntsX(ctx context.Context) []int {
	v, err := hs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hs *HexSelect) IntX(ctx context.Context) int {
	v, err := hs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HexSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hs *HexSelect) Float64sX(ctx context.Context) []float64 {
	v, err := hs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hs *HexSelect) Float64X(ctx context.Context) float64 {
	v, err := hs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HexSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hs *HexSelect) BoolsX(ctx context.Context) []bool {
	v, err := hs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (hs *HexSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{hex.Label}
	default:
		err = fmt.Errorf("ent: HexSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hs *HexSelect) BoolX(ctx context.Context) bool {
	v, err := hs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hs *HexSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sqlQuery().Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hs *HexSelect) sqlQuery() sql.Querier {
	selector := hs.sql
	selector.Select(selector.Columns(hs.fields...)...)
	return selector
}