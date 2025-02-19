// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/dependsonskipped"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DependsOnSkippedDelete is the builder for deleting a DependsOnSkipped entity.
type DependsOnSkippedDelete struct {
	config
	hooks    []Hook
	mutation *DependsOnSkippedMutation
}

// Where appends a list predicates to the DependsOnSkippedDelete builder.
func (dosd *DependsOnSkippedDelete) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedDelete {
	dosd.mutation.Where(ps...)
	return dosd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dosd *DependsOnSkippedDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dosd.hooks) == 0 {
		affected, err = dosd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DependsOnSkippedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dosd.mutation = mutation
			affected, err = dosd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dosd.hooks) - 1; i >= 0; i-- {
			if dosd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dosd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dosd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dosd *DependsOnSkippedDelete) ExecX(ctx context.Context) int {
	n, err := dosd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dosd *DependsOnSkippedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dependsonskipped.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dependsonskipped.FieldID,
			},
		},
	}
	if ps := dosd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dosd.driver, _spec)
}

// DependsOnSkippedDeleteOne is the builder for deleting a single DependsOnSkipped entity.
type DependsOnSkippedDeleteOne struct {
	dosd *DependsOnSkippedDelete
}

// Exec executes the deletion query.
func (dosdo *DependsOnSkippedDeleteOne) Exec(ctx context.Context) error {
	n, err := dosdo.dosd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dependsonskipped.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dosdo *DependsOnSkippedDeleteOne) ExecX(ctx context.Context) {
	dosdo.dosd.ExecX(ctx)
}
