// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/juicycleff/frank/ent/featureflag"
	"github.com/juicycleff/frank/ent/predicate"
)

// FeatureFlagDelete is the builder for deleting a FeatureFlag entity.
type FeatureFlagDelete struct {
	config
	hooks    []Hook
	mutation *FeatureFlagMutation
}

// Where appends a list predicates to the FeatureFlagDelete builder.
func (ffd *FeatureFlagDelete) Where(ps ...predicate.FeatureFlag) *FeatureFlagDelete {
	ffd.mutation.Where(ps...)
	return ffd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ffd *FeatureFlagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ffd.sqlExec, ffd.mutation, ffd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ffd *FeatureFlagDelete) ExecX(ctx context.Context) int {
	n, err := ffd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ffd *FeatureFlagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(featureflag.Table, sqlgraph.NewFieldSpec(featureflag.FieldID, field.TypeString))
	if ps := ffd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ffd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ffd.mutation.done = true
	return affected, err
}

// FeatureFlagDeleteOne is the builder for deleting a single FeatureFlag entity.
type FeatureFlagDeleteOne struct {
	ffd *FeatureFlagDelete
}

// Where appends a list predicates to the FeatureFlagDelete builder.
func (ffdo *FeatureFlagDeleteOne) Where(ps ...predicate.FeatureFlag) *FeatureFlagDeleteOne {
	ffdo.ffd.mutation.Where(ps...)
	return ffdo
}

// Exec executes the deletion query.
func (ffdo *FeatureFlagDeleteOne) Exec(ctx context.Context) error {
	n, err := ffdo.ffd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{featureflag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ffdo *FeatureFlagDeleteOne) ExecX(ctx context.Context) {
	if err := ffdo.Exec(ctx); err != nil {
		panic(err)
	}
}
