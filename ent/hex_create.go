// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/hex"
	"github.com/andrewmthomas87/catan/ent/numbertoken"
	"github.com/andrewmthomas87/catan/ent/robber"
	"github.com/andrewmthomas87/catan/ent/settlement"
)

// HexCreate is the builder for creating a Hex entity.
type HexCreate struct {
	config
	mutation *HexMutation
	hooks    []Hook
}

// SetX sets the "x" field.
func (hc *HexCreate) SetX(i int) *HexCreate {
	hc.mutation.SetX(i)
	return hc
}

// SetY sets the "y" field.
func (hc *HexCreate) SetY(i int) *HexCreate {
	hc.mutation.SetY(i)
	return hc
}

// SetTerrain sets the "terrain" field.
func (hc *HexCreate) SetTerrain(h hex.Terrain) *HexCreate {
	hc.mutation.SetTerrain(h)
	return hc
}

// SetNumberTokenID sets the "number_token" edge to the NumberToken entity by ID.
func (hc *HexCreate) SetNumberTokenID(id int) *HexCreate {
	hc.mutation.SetNumberTokenID(id)
	return hc
}

// SetNillableNumberTokenID sets the "number_token" edge to the NumberToken entity by ID if the given value is not nil.
func (hc *HexCreate) SetNillableNumberTokenID(id *int) *HexCreate {
	if id != nil {
		hc = hc.SetNumberTokenID(*id)
	}
	return hc
}

// SetNumberToken sets the "number_token" edge to the NumberToken entity.
func (hc *HexCreate) SetNumberToken(n *NumberToken) *HexCreate {
	return hc.SetNumberTokenID(n.ID)
}

// SetRobberID sets the "robber" edge to the Robber entity by ID.
func (hc *HexCreate) SetRobberID(id int) *HexCreate {
	hc.mutation.SetRobberID(id)
	return hc
}

// SetNillableRobberID sets the "robber" edge to the Robber entity by ID if the given value is not nil.
func (hc *HexCreate) SetNillableRobberID(id *int) *HexCreate {
	if id != nil {
		hc = hc.SetRobberID(*id)
	}
	return hc
}

// SetRobber sets the "robber" edge to the Robber entity.
func (hc *HexCreate) SetRobber(r *Robber) *HexCreate {
	return hc.SetRobberID(r.ID)
}

// AddSettlementIDs adds the "settlements" edge to the Settlement entity by IDs.
func (hc *HexCreate) AddSettlementIDs(ids ...int) *HexCreate {
	hc.mutation.AddSettlementIDs(ids...)
	return hc
}

// AddSettlements adds the "settlements" edges to the Settlement entity.
func (hc *HexCreate) AddSettlements(s ...*Settlement) *HexCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return hc.AddSettlementIDs(ids...)
}

// Mutation returns the HexMutation object of the builder.
func (hc *HexCreate) Mutation() *HexMutation {
	return hc.mutation
}

// Save creates the Hex in the database.
func (hc *HexCreate) Save(ctx context.Context) (*Hex, error) {
	var (
		err  error
		node *Hex
	)
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HexMutation)
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
func (hc *HexCreate) SaveX(ctx context.Context) *Hex {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (hc *HexCreate) check() error {
	if _, ok := hc.mutation.X(); !ok {
		return &ValidationError{Name: "x", err: errors.New("ent: missing required field \"x\"")}
	}
	if _, ok := hc.mutation.Y(); !ok {
		return &ValidationError{Name: "y", err: errors.New("ent: missing required field \"y\"")}
	}
	if _, ok := hc.mutation.Terrain(); !ok {
		return &ValidationError{Name: "terrain", err: errors.New("ent: missing required field \"terrain\"")}
	}
	if v, ok := hc.mutation.Terrain(); ok {
		if err := hex.TerrainValidator(v); err != nil {
			return &ValidationError{Name: "terrain", err: fmt.Errorf("ent: validator failed for field \"terrain\": %w", err)}
		}
	}
	return nil
}

func (hc *HexCreate) sqlSave(ctx context.Context) (*Hex, error) {
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

func (hc *HexCreate) createSpec() (*Hex, *sqlgraph.CreateSpec) {
	var (
		_node = &Hex{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hex.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hex.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.X(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldX,
		})
		_node.X = value
	}
	if value, ok := hc.mutation.Y(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: hex.FieldY,
		})
		_node.Y = value
	}
	if value, ok := hc.mutation.Terrain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: hex.FieldTerrain,
		})
		_node.Terrain = value
	}
	if nodes := hc.mutation.NumberTokenIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.RobberIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.SettlementsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HexCreateBulk is the builder for creating many Hex entities in bulk.
type HexCreateBulk struct {
	config
	builders []*HexCreate
}

// Save creates the Hex entities in the database.
func (hcb *HexCreateBulk) Save(ctx context.Context) ([]*Hex, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hex, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HexMutation)
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
func (hcb *HexCreateBulk) SaveX(ctx context.Context) []*Hex {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
