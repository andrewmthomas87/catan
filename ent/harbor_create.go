// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/harbor"
	"github.com/andrewmthomas87/catan/ent/settlement"
)

// HarborCreate is the builder for creating a Harbor entity.
type HarborCreate struct {
	config
	mutation *HarborMutation
	hooks    []Hook
}

// SetX sets the "x" field.
func (hc *HarborCreate) SetX(i int) *HarborCreate {
	hc.mutation.SetX(i)
	return hc
}

// SetY sets the "y" field.
func (hc *HarborCreate) SetY(i int) *HarborCreate {
	hc.mutation.SetY(i)
	return hc
}

// SetResource sets the "resource" field.
func (hc *HarborCreate) SetResource(h harbor.Resource) *HarborCreate {
	hc.mutation.SetResource(h)
	return hc
}

// SetSettlementID sets the "settlement" edge to the Settlement entity by ID.
func (hc *HarborCreate) SetSettlementID(id int) *HarborCreate {
	hc.mutation.SetSettlementID(id)
	return hc
}

// SetNillableSettlementID sets the "settlement" edge to the Settlement entity by ID if the given value is not nil.
func (hc *HarborCreate) SetNillableSettlementID(id *int) *HarborCreate {
	if id != nil {
		hc = hc.SetSettlementID(*id)
	}
	return hc
}

// SetSettlement sets the "settlement" edge to the Settlement entity.
func (hc *HarborCreate) SetSettlement(s *Settlement) *HarborCreate {
	return hc.SetSettlementID(s.ID)
}

// Mutation returns the HarborMutation object of the builder.
func (hc *HarborCreate) Mutation() *HarborMutation {
	return hc.mutation
}

// Save creates the Harbor in the database.
func (hc *HarborCreate) Save(ctx context.Context) (*Harbor, error) {
	var (
		err  error
		node *Harbor
	)
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HarborMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			node, err = hc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HarborCreate) SaveX(ctx context.Context) *Harbor {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (hc *HarborCreate) check() error {
	if _, ok := hc.mutation.X(); !ok {
		return &ValidationError{Name: "x", err: errors.New("ent: missing required field \"x\"")}
	}
	if _, ok := hc.mutation.Y(); !ok {
		return &ValidationError{Name: "y", err: errors.New("ent: missing required field \"y\"")}
	}
	if _, ok := hc.mutation.Resource(); !ok {
		return &ValidationError{Name: "resource", err: errors.New("ent: missing required field \"resource\"")}
	}
	if v, ok := hc.mutation.Resource(); ok {
		if err := harbor.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf("ent: validator failed for field \"resource\": %w", err)}
		}
	}
	return nil
}

func (hc *HarborCreate) sqlSave(ctx context.Context) (*Harbor, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HarborCreate) createSpec() (*Harbor, *sqlgraph.CreateSpec) {
	var (
		_node = &Harbor{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: harbor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: harbor.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.X(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: harbor.FieldX,
		})
		_node.X = value
	}
	if value, ok := hc.mutation.Y(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: harbor.FieldY,
		})
		_node.Y = value
	}
	if value, ok := hc.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: harbor.FieldResource,
		})
		_node.Resource = value
	}
	if nodes := hc.mutation.SettlementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   harbor.SettlementTable,
			Columns: []string{harbor.SettlementColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HarborCreateBulk is the builder for creating many Harbor entities in bulk.
type HarborCreateBulk struct {
	config
	builders []*HarborCreate
}

// Save creates the Harbor entities in the database.
func (hcb *HarborCreateBulk) Save(ctx context.Context) ([]*Harbor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Harbor, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HarborMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HarborCreateBulk) SaveX(ctx context.Context) []*Harbor {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}