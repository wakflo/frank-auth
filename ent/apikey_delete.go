// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/juicycleff/frank/ent/apikey"
	"github.com/juicycleff/frank/ent/predicate"
)

// ApiKeyDelete is the builder for deleting a ApiKey entity.
type ApiKeyDelete struct {
	config
	hooks    []Hook
	mutation *ApiKeyMutation
}

// Where appends a list predicates to the ApiKeyDelete builder.
func (akd *ApiKeyDelete) Where(ps ...predicate.ApiKey) *ApiKeyDelete {
	akd.mutation.Where(ps...)
	return akd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (akd *ApiKeyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, akd.sqlExec, akd.mutation, akd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (akd *ApiKeyDelete) ExecX(ctx context.Context) int {
	n, err := akd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (akd *ApiKeyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(apikey.Table, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeString))
	if ps := akd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, akd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	akd.mutation.done = true
	return affected, err
}

// ApiKeyDeleteOne is the builder for deleting a single ApiKey entity.
type ApiKeyDeleteOne struct {
	akd *ApiKeyDelete
}

// Where appends a list predicates to the ApiKeyDelete builder.
func (akdo *ApiKeyDeleteOne) Where(ps ...predicate.ApiKey) *ApiKeyDeleteOne {
	akdo.akd.mutation.Where(ps...)
	return akdo
}

// Exec executes the deletion query.
func (akdo *ApiKeyDeleteOne) Exec(ctx context.Context) error {
	n, err := akdo.akd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{apikey.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (akdo *ApiKeyDeleteOne) ExecX(ctx context.Context) {
	if err := akdo.Exec(ctx); err != nil {
		panic(err)
	}
}
