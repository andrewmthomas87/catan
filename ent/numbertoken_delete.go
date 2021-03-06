// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andrewmthomas87/catan/ent/numbertoken"
	"github.com/andrewmthomas87/catan/ent/predicate"
)

// NumberTokenDelete is the builder for deleting a NumberToken entity.
type NumberTokenDelete struct {
	config
	hooks    []Hook
	mutation *NumberTokenMutation
}

// Where adds a new predicate to the NumberTokenDelete builder.
func (ntd *NumberTokenDelete) Where(ps ...predicate.NumberToken) *NumberTokenDelete {
	ntd.mutation.predicates = append(ntd.mutation.predicates, ps...)
	return ntd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ntd *NumberTokenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ntd.hooks) == 0 {
		affected, err = ntd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NumberTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ntd.mutation = mutation
			affected, err = ntd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ntd.hooks) - 1; i >= 0; i-- {
			mut = ntd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntd *NumberTokenDelete) ExecX(ctx context.Context) int {
	n, err := ntd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ntd *NumberTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: numbertoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: numbertoken.FieldID,
			},
		},
	}
	if ps := ntd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ntd.driver, _spec)
}

// NumberTokenDeleteOne is the builder for deleting a single NumberToken entity.
type NumberTokenDeleteOne struct {
	ntd *NumberTokenDelete
}

// Exec executes the deletion query.
func (ntdo *NumberTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := ntdo.ntd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{numbertoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ntdo *NumberTokenDeleteOne) ExecX(ctx context.Context) {
	ntdo.ntd.ExecX(ctx)
}
