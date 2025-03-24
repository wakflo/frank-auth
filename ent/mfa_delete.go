// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/juicycleff/frank/ent/mfa"
	"github.com/juicycleff/frank/ent/predicate"
)

// MFADelete is the builder for deleting a MFA entity.
type MFADelete struct {
	config
	hooks    []Hook
	mutation *MFAMutation
}

// Where appends a list predicates to the MFADelete builder.
func (md *MFADelete) Where(ps ...predicate.MFA) *MFADelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MFADelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MFADelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MFADelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(mfa.Table, sqlgraph.NewFieldSpec(mfa.FieldID, field.TypeString))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MFADeleteOne is the builder for deleting a single MFA entity.
type MFADeleteOne struct {
	md *MFADelete
}

// Where appends a list predicates to the MFADelete builder.
func (mdo *MFADeleteOne) Where(ps ...predicate.MFA) *MFADeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MFADeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mfa.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MFADeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
