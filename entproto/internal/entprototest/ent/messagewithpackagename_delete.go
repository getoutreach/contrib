// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithpackagename"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithPackageNameDelete is the builder for deleting a MessageWithPackageName entity.
type MessageWithPackageNameDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithPackageNameMutation
}

// Where appends a list predicates to the MessageWithPackageNameDelete builder.
func (mwpnd *MessageWithPackageNameDelete) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameDelete {
	mwpnd.mutation.Where(ps...)
	return mwpnd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwpnd *MessageWithPackageNameDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mwpnd.hooks) == 0 {
		affected, err = mwpnd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithPackageNameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwpnd.mutation = mutation
			affected, err = mwpnd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mwpnd.hooks) - 1; i >= 0; i-- {
			if mwpnd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwpnd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwpnd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwpnd *MessageWithPackageNameDelete) ExecX(ctx context.Context) int {
	n, err := mwpnd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwpnd *MessageWithPackageNameDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: messagewithpackagename.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithpackagename.FieldID,
			},
		},
	}
	if ps := mwpnd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mwpnd.driver, _spec)
}

// MessageWithPackageNameDeleteOne is the builder for deleting a single MessageWithPackageName entity.
type MessageWithPackageNameDeleteOne struct {
	mwpnd *MessageWithPackageNameDelete
}

// Exec executes the deletion query.
func (mwpndo *MessageWithPackageNameDeleteOne) Exec(ctx context.Context) error {
	n, err := mwpndo.mwpnd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithpackagename.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwpndo *MessageWithPackageNameDeleteOne) ExecX(ctx context.Context) {
	mwpndo.mwpnd.ExecX(ctx)
}
