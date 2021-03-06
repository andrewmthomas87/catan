// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/hex"
	"github.com/andrewmthomas87/catan/ent/numbertoken"
	"github.com/andrewmthomas87/catan/ent/predicate"
	"github.com/andrewmthomas87/catan/ent/robber"
	"github.com/andrewmthomas87/catan/ent/settlement"
)

// HexUpdate is the builder for updating Hex entities.
type HexUpdate struct {
	config
	hooks    []Hook
	mutation *HexMutation
}

// Where adds a new predicate for the HexUpdate builder.
func (hu *HexUpdate) Where(ps ...predicate.Hex) *HexUpdate {
	hu.mutation.predicates = append(hu.mutation.predicates, ps...)
	return hu
}

// SetX sets the "x" field.
func (hu *HexUpdate) SetX(i int) *HexUpdate {
	hu.mutation.ResetX()
	hu.mutation.SetX(i)
	return hu
}

// AddX adds i to the "x" field.
func (hu *HexUpdate) AddX(i int) *HexUpdate {
	hu.mutation.AddX(i)
	return hu
}

// SetY sets the "y" field.
func (hu *HexUpdate) SetY(i int) *HexUpdate {
	hu.mutation.ResetY()
	hu.mutation.SetY(i)
	return hu
}

// AddY adds i to the "y" field.
func (hu *HexUpdate) AddY(i int) *HexUpdate {
	hu.mutation.AddY(i)
	return hu
}

// SetTerrain sets the "terrain" field.
func (hu *HexUpdate) SetTerrain(h hex.Terrain) *HexUpdate {
	hu.mutation.SetTerrain(h)
	return hu
}

// SetNumberTokenID sets the "number_token" edge to the NumberToken entity by ID.
func (hu *HexUpdate) SetNumberTokenID(id int) *HexUpdate {
	hu.mutation.SetNumberTokenID(id)
	return hu
}

// SetNillableNumberTokenID sets the "number_token" edge to the NumberToken entity by ID if the given value is not nil.
func (hu *HexUpdate) SetNillableNumberTokenID(id *int) *HexUpdate {
	if id != nil {
		hu = hu.SetNumberTokenID(*id)
	}
	return hu
}

// SetNumberToken sets the "number_token" edge to the NumberToken entity.
func (hu *HexUpdate) SetNumberToken(n *NumberToken) *HexUpdate {
	return hu.SetNumberTokenID(n.ID)
}

// SetRobberID sets the "robber" edge to the Robber entity by ID.
func (hu *HexUpdate) SetRobberID(id int) *HexUpdate {
	hu.mutation.SetRobberID(id)
	return hu
}

// SetNillableRobberID sets the "robber" edge to the Robber entity by ID if the given value is not nil.
func (hu *HexUpdate) SetNillableRobberID(id *int) *HexUpdate {
	if id != nil {
		hu = hu.SetRobberID(*id)
	}
	return hu
}

// SetRobber sets the "robber" edge to the Robber entity.
func (hu *HexUpdate) SetRobber(r *Robber) *HexUpdate {
	return hu.SetRobberID(r.ID)
}

// AddSettlementIDs adds the "settlements" edge to the Settlement entity by IDs.
func (hu *HexUpdate) AddSettlementIDs(ids ...int) *HexUpdate {
	hu.mutation.AddSettlementIDs(ids...)
	return hu
}

// AddSettlements adds the "settlements" edges to the Settlement entity.
func (hu *HexUpdate) AddSettlements(s ...*Settlement) *HexUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return hu.AddSettlementIDs(ids...)
}

// Mutation returns the HexMutation object of the builder.
func (hu *HexUpdate) Mutation() *HexMutation {
	return hu.mutation
}

// ClearNumberToken clears the "number_token" edge to the NumberToken entity.
func (hu *HexUpdate) ClearNumberToken() *HexUpdate {
	hu.mutation.ClearNumberToken()
	return hu
}

// ClearRobber clears the "robber" edge to the Robber entity.
func (hu *HexUpdate) ClearRobber() *HexUpdate {
	hu.mutation.ClearRobber()
	return hu
}

// ClearSettlements clears all "settlements" edges to the Settlement entity.
func (hu *HexUpdate) ClearSettlements() *HexUpdate {
	hu.mutation.ClearSettlements()
	return hu
}

// RemoveSettlementIDs removes the "settlements" edge to Settlement entities by IDs.
func (hu *HexUpdate) RemoveSettlementIDs(ids ...int) *HexUpdate {
	hu.mutation.RemoveSettlementIDs(ids...)
	return hu
}

// RemoveSettlements removes "settlements" edges to Settlement entities.
func (hu *HexUpdate) RemoveSettlements(s ...*Settlement) *HexUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return hu.RemoveSettlementIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HexUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hu.hooks) == 0 {
		if err = hu.check(); err != nil {
			return 0, err
		}
		affected, err = hu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hu.check(); err != nil {
				return 0, err
			}
			hu.mutation = mutation
			affected, err = hu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hu.hooks) - 1; i >= 0; i-- {
			mut = hu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HexUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HexUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HexUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HexUpdate) check() error {
	if v, ok := hu.mutation.Terrain(); ok {
		if err := hex.TerrainValidator(v); err != nil {
			return &ValidationError{Name: "terrain", err: fmt.Errorf("ent: validator failed for field \"terrain\": %w", err)}
		}
	}
	return nil
}

func (hu *HexUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hex.Table,
			Columns: hex.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hex.FieldID,
			},
		},
	}
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.X(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldX,
		})
	}
	if value, ok := hu.mutation.AddedX(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldX,
		})
	}
	if value, ok := hu.mutation.Y(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldY,
		})
	}
	if value, ok := hu.mutation.AddedY(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldY,
		})
	}
	if value, ok := hu.mutation.Terrain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: hex.FieldTerrain,
		})
	}
	if hu.mutation.NumberTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.NumberTokenTable,
			Columns: []string{hex.NumberTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: numbertoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.NumberTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.NumberTokenTable,
			Columns: []string{hex.NumberTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: numbertoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.RobberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.RobberTable,
			Columns: []string{hex.RobberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: robber.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RobberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.RobberTable,
			Columns: []string{hex.RobberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: robber.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.SettlementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hex.SettlementsTable,
			Columns: hex.SettlementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settlement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedSettlementsIDs(); len(nodes) > 0 && !hu.mutation.SettlementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hex.SettlementsTable,
			Columns: hex.SettlementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settlement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.SettlementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hex.SettlementsTable,
			Columns: hex.SettlementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settlement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hex.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HexUpdateOne is the builder for updating a single Hex entity.
type HexUpdateOne struct {
	config
	hooks    []Hook
	mutation *HexMutation
}

// SetX sets the "x" field.
func (huo *HexUpdateOne) SetX(i int) *HexUpdateOne {
	huo.mutation.ResetX()
	huo.mutation.SetX(i)
	return huo
}

// AddX adds i to the "x" field.
func (huo *HexUpdateOne) AddX(i int) *HexUpdateOne {
	huo.mutation.AddX(i)
	return huo
}

// SetY sets the "y" field.
func (huo *HexUpdateOne) SetY(i int) *HexUpdateOne {
	huo.mutation.ResetY()
	huo.mutation.SetY(i)
	return huo
}

// AddY adds i to the "y" field.
func (huo *HexUpdateOne) AddY(i int) *HexUpdateOne {
	huo.mutation.AddY(i)
	return huo
}

// SetTerrain sets the "terrain" field.
func (huo *HexUpdateOne) SetTerrain(h hex.Terrain) *HexUpdateOne {
	huo.mutation.SetTerrain(h)
	return huo
}

// SetNumberTokenID sets the "number_token" edge to the NumberToken entity by ID.
func (huo *HexUpdateOne) SetNumberTokenID(id int) *HexUpdateOne {
	huo.mutation.SetNumberTokenID(id)
	return huo
}

// SetNillableNumberTokenID sets the "number_token" edge to the NumberToken entity by ID if the given value is not nil.
func (huo *HexUpdateOne) SetNillableNumberTokenID(id *int) *HexUpdateOne {
	if id != nil {
		huo = huo.SetNumberTokenID(*id)
	}
	return huo
}

// SetNumberToken sets the "number_token" edge to the NumberToken entity.
func (huo *HexUpdateOne) SetNumberToken(n *NumberToken) *HexUpdateOne {
	return huo.SetNumberTokenID(n.ID)
}

// SetRobberID sets the "robber" edge to the Robber entity by ID.
func (huo *HexUpdateOne) SetRobberID(id int) *HexUpdateOne {
	huo.mutation.SetRobberID(id)
	return huo
}

// SetNillableRobberID sets the "robber" edge to the Robber entity by ID if the given value is not nil.
func (huo *HexUpdateOne) SetNillableRobberID(id *int) *HexUpdateOne {
	if id != nil {
		huo = huo.SetRobberID(*id)
	}
	return huo
}

// SetRobber sets the "robber" edge to the Robber entity.
func (huo *HexUpdateOne) SetRobber(r *Robber) *HexUpdateOne {
	return huo.SetRobberID(r.ID)
}

// AddSettlementIDs adds the "settlements" edge to the Settlement entity by IDs.
func (huo *HexUpdateOne) AddSettlementIDs(ids ...int) *HexUpdateOne {
	huo.mutation.AddSettlementIDs(ids...)
	return huo
}

// AddSettlements adds the "settlements" edges to the Settlement entity.
func (huo *HexUpdateOne) AddSettlements(s ...*Settlement) *HexUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return huo.AddSettlementIDs(ids...)
}

// Mutation returns the HexMutation object of the builder.
func (huo *HexUpdateOne) Mutation() *HexMutation {
	return huo.mutation
}

// ClearNumberToken clears the "number_token" edge to the NumberToken entity.
func (huo *HexUpdateOne) ClearNumberToken() *HexUpdateOne {
	huo.mutation.ClearNumberToken()
	return huo
}

// ClearRobber clears the "robber" edge to the Robber entity.
func (huo *HexUpdateOne) ClearRobber() *HexUpdateOne {
	huo.mutation.ClearRobber()
	return huo
}

// ClearSettlements clears all "settlements" edges to the Settlement entity.
func (huo *HexUpdateOne) ClearSettlements() *HexUpdateOne {
	huo.mutation.ClearSettlements()
	return huo
}

// RemoveSettlementIDs removes the "settlements" edge to Settlement entities by IDs.
func (huo *HexUpdateOne) RemoveSettlementIDs(ids ...int) *HexUpdateOne {
	huo.mutation.RemoveSettlementIDs(ids...)
	return huo
}

// RemoveSettlements removes "settlements" edges to Settlement entities.
func (huo *HexUpdateOne) RemoveSettlements(s ...*Settlement) *HexUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return huo.RemoveSettlementIDs(ids...)
}

// Save executes the query and returns the updated Hex entity.
func (huo *HexUpdateOne) Save(ctx context.Context) (*Hex, error) {
	var (
		err  error
		node *Hex
	)
	if len(huo.hooks) == 0 {
		if err = huo.check(); err != nil {
			return nil, err
		}
		node, err = huo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = huo.check(); err != nil {
				return nil, err
			}
			huo.mutation = mutation
			node, err = huo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(huo.hooks) - 1; i >= 0; i-- {
			mut = huo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, huo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HexUpdateOne) SaveX(ctx context.Context) *Hex {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HexUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HexUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HexUpdateOne) check() error {
	if v, ok := huo.mutation.Terrain(); ok {
		if err := hex.TerrainValidator(v); err != nil {
			return &ValidationError{Name: "terrain", err: fmt.Errorf("ent: validator failed for field \"terrain\": %w", err)}
		}
	}
	return nil
}

func (huo *HexUpdateOne) sqlSave(ctx context.Context) (_node *Hex, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hex.Table,
			Columns: hex.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hex.FieldID,
			},
		},
	}
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Hex.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.X(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldX,
		})
	}
	if value, ok := huo.mutation.AddedX(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldX,
		})
	}
	if value, ok := huo.mutation.Y(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldY,
		})
	}
	if value, ok := huo.mutation.AddedY(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldY,
		})
	}
	if value, ok := huo.mutation.Terrain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: hex.FieldTerrain,
		})
	}
	if huo.mutation.NumberTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.NumberTokenTable,
			Columns: []string{hex.NumberTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: numbertoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.NumberTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.NumberTokenTable,
			Columns: []string{hex.NumberTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: numbertoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.RobberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.RobberTable,
			Columns: []string{hex.RobberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: robber.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RobberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hex.RobberTable,
			Columns: []string{hex.RobberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: robber.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.SettlementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hex.SettlementsTable,
			Columns: hex.SettlementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settlement.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedSettlementsIDs(); len(nodes) > 0 && !huo.mutation.SettlementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hex.SettlementsTable,
			Columns: hex.SettlementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settlement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.SettlementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hex.SettlementsTable,
			Columns: hex.SettlementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settlement.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Hex{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hex.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
