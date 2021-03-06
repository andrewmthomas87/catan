// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/hex"
	"github.com/andrewmthomas87/catan/ent/predicate"
	"github.com/andrewmthomas87/catan/ent/robber"
)

// RobberUpdate is the builder for updating Robber entities.
type RobberUpdate struct {
	config
	hooks    []Hook
	mutation *RobberMutation
}

// Where adds a new predicate for the RobberUpdate builder.
func (ru *RobberUpdate) Where(ps ...predicate.Robber) *RobberUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetHexID sets the "hex" edge to the Hex entity by ID.
func (ru *RobberUpdate) SetHexID(id int) *RobberUpdate {
	ru.mutation.SetHexID(id)
	return ru
}

// SetNillableHexID sets the "hex" edge to the Hex entity by ID if the given value is not nil.
func (ru *RobberUpdate) SetNillableHexID(id *int) *RobberUpdate {
	if id != nil {
		ru = ru.SetHexID(*id)
	}
	return ru
}

// SetHex sets the "hex" edge to the Hex entity.
func (ru *RobberUpdate) SetHex(h *Hex) *RobberUpdate {
	return ru.SetHexID(h.ID)
}

// Mutation returns the RobberMutation object of the builder.
func (ru *RobberUpdate) Mutation() *RobberMutation {
	return ru.mutation
}

// ClearHex clears the "hex" edge to the Hex entity.
func (ru *RobberUpdate) ClearHex() *RobberUpdate {
	ru.mutation.ClearHex()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RobberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RobberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RobberUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RobberUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RobberUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RobberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   robber.Table,
			Columns: robber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: robber.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ru.mutation.HexCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   robber.HexTable,
			Columns: []string{robber.HexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hex.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.HexIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   robber.HexTable,
			Columns: []string{robber.HexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hex.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{robber.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RobberUpdateOne is the builder for updating a single Robber entity.
type RobberUpdateOne struct {
	config
	hooks    []Hook
	mutation *RobberMutation
}

// SetHexID sets the "hex" edge to the Hex entity by ID.
func (ruo *RobberUpdateOne) SetHexID(id int) *RobberUpdateOne {
	ruo.mutation.SetHexID(id)
	return ruo
}

// SetNillableHexID sets the "hex" edge to the Hex entity by ID if the given value is not nil.
func (ruo *RobberUpdateOne) SetNillableHexID(id *int) *RobberUpdateOne {
	if id != nil {
		ruo = ruo.SetHexID(*id)
	}
	return ruo
}

// SetHex sets the "hex" edge to the Hex entity.
func (ruo *RobberUpdateOne) SetHex(h *Hex) *RobberUpdateOne {
	return ruo.SetHexID(h.ID)
}

// Mutation returns the RobberMutation object of the builder.
func (ruo *RobberUpdateOne) Mutation() *RobberMutation {
	return ruo.mutation
}

// ClearHex clears the "hex" edge to the Hex entity.
func (ruo *RobberUpdateOne) ClearHex() *RobberUpdateOne {
	ruo.mutation.ClearHex()
	return ruo
}

// Save executes the query and returns the updated Robber entity.
func (ruo *RobberUpdateOne) Save(ctx context.Context) (*Robber, error) {
	var (
		err  error
		node *Robber
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RobberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RobberUpdateOne) SaveX(ctx context.Context) *Robber {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RobberUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RobberUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RobberUpdateOne) sqlSave(ctx context.Context) (_node *Robber, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   robber.Table,
			Columns: robber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: robber.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Robber.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ruo.mutation.HexCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   robber.HexTable,
			Columns: []string{robber.HexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hex.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.HexIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   robber.HexTable,
			Columns: []string{robber.HexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hex.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Robber{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{robber.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
