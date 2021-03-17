// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/road"
)

// RoadCreate is the builder for creating a Road entity.
type RoadCreate struct {
	config
	mutation *RoadMutation
	hooks    []Hook
}

// SetX sets the "x" field.
func (rc *RoadCreate) SetX(i int) *RoadCreate {
	rc.mutation.SetX(i)
	return rc
}

// SetY sets the "y" field.
func (rc *RoadCreate) SetY(i int) *RoadCreate {
	rc.mutation.SetY(i)
	return rc
}

// Mutation returns the RoadMutation object of the builder.
func (rc *RoadCreate) Mutation() *RoadMutation {
	return rc.mutation
}

// Save creates the Road in the database.
func (rc *RoadCreate) Save(ctx context.Context) (*Road, error) {
	var (
		err  error
		node *Road
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoadCreate) SaveX(ctx context.Context) *Road {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoadCreate) check() error {
	if _, ok := rc.mutation.X(); !ok {
		return &ValidationError{Name: "x", err: errors.New("ent: missing required field \"x\"")}
	}
	if _, ok := rc.mutation.Y(); !ok {
		return &ValidationError{Name: "y", err: errors.New("ent: missing required field \"y\"")}
	}
	return nil
}

func (rc *RoadCreate) sqlSave(ctx context.Context) (*Road, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RoadCreate) createSpec() (*Road, *sqlgraph.CreateSpec) {
	var (
		_node = &Road{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: road.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: road.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.X(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: road.FieldX,
		})
		_node.X = value
	}
	if value, ok := rc.mutation.Y(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: road.FieldY,
		})
		_node.Y = value
	}
	return _node, _spec
}

// RoadCreateBulk is the builder for creating many Road entities in bulk.
type RoadCreateBulk struct {
	config
	builders []*RoadCreate
}

// Save creates the Road entities in the database.
func (rcb *RoadCreateBulk) Save(ctx context.Context) ([]*Road, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Road, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoadMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoadCreateBulk) SaveX(ctx context.Context) []*Road {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}