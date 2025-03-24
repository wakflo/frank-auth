// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/juicycleff/frank/ent/predicate"
	"github.com/juicycleff/frank/ent/ssostate"
)

// SSOStateUpdate is the builder for updating SSOState entities.
type SSOStateUpdate struct {
	config
	hooks    []Hook
	mutation *SSOStateMutation
}

// Where appends a list predicates to the SSOStateUpdate builder.
func (ssu *SSOStateUpdate) Where(ps ...predicate.SSOState) *SSOStateUpdate {
	ssu.mutation.Where(ps...)
	return ssu
}

// SetUpdatedAt sets the "updated_at" field.
func (ssu *SSOStateUpdate) SetUpdatedAt(t time.Time) *SSOStateUpdate {
	ssu.mutation.SetUpdatedAt(t)
	return ssu
}

// SetState sets the "state" field.
func (ssu *SSOStateUpdate) SetState(s string) *SSOStateUpdate {
	ssu.mutation.SetState(s)
	return ssu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ssu *SSOStateUpdate) SetNillableState(s *string) *SSOStateUpdate {
	if s != nil {
		ssu.SetState(*s)
	}
	return ssu
}

// SetData sets the "data" field.
func (ssu *SSOStateUpdate) SetData(s string) *SSOStateUpdate {
	ssu.mutation.SetData(s)
	return ssu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (ssu *SSOStateUpdate) SetNillableData(s *string) *SSOStateUpdate {
	if s != nil {
		ssu.SetData(*s)
	}
	return ssu
}

// SetExpiresAt sets the "expires_at" field.
func (ssu *SSOStateUpdate) SetExpiresAt(t time.Time) *SSOStateUpdate {
	ssu.mutation.SetExpiresAt(t)
	return ssu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (ssu *SSOStateUpdate) SetNillableExpiresAt(t *time.Time) *SSOStateUpdate {
	if t != nil {
		ssu.SetExpiresAt(*t)
	}
	return ssu
}

// Mutation returns the SSOStateMutation object of the builder.
func (ssu *SSOStateUpdate) Mutation() *SSOStateMutation {
	return ssu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ssu *SSOStateUpdate) Save(ctx context.Context) (int, error) {
	ssu.defaults()
	return withHooks(ctx, ssu.sqlSave, ssu.mutation, ssu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssu *SSOStateUpdate) SaveX(ctx context.Context) int {
	affected, err := ssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ssu *SSOStateUpdate) Exec(ctx context.Context) error {
	_, err := ssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssu *SSOStateUpdate) ExecX(ctx context.Context) {
	if err := ssu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ssu *SSOStateUpdate) defaults() {
	if _, ok := ssu.mutation.UpdatedAt(); !ok {
		v := ssostate.UpdateDefaultUpdatedAt()
		ssu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ssu *SSOStateUpdate) check() error {
	if v, ok := ssu.mutation.Data(); ok {
		if err := ssostate.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "SSOState.data": %w`, err)}
		}
	}
	return nil
}

func (ssu *SSOStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ssu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ssostate.Table, ssostate.Columns, sqlgraph.NewFieldSpec(ssostate.FieldID, field.TypeString))
	if ps := ssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ssu.mutation.UpdatedAt(); ok {
		_spec.SetField(ssostate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ssu.mutation.State(); ok {
		_spec.SetField(ssostate.FieldState, field.TypeString, value)
	}
	if value, ok := ssu.mutation.Data(); ok {
		_spec.SetField(ssostate.FieldData, field.TypeString, value)
	}
	if value, ok := ssu.mutation.ExpiresAt(); ok {
		_spec.SetField(ssostate.FieldExpiresAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ssostate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ssu.mutation.done = true
	return n, nil
}

// SSOStateUpdateOne is the builder for updating a single SSOState entity.
type SSOStateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SSOStateMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ssuo *SSOStateUpdateOne) SetUpdatedAt(t time.Time) *SSOStateUpdateOne {
	ssuo.mutation.SetUpdatedAt(t)
	return ssuo
}

// SetState sets the "state" field.
func (ssuo *SSOStateUpdateOne) SetState(s string) *SSOStateUpdateOne {
	ssuo.mutation.SetState(s)
	return ssuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ssuo *SSOStateUpdateOne) SetNillableState(s *string) *SSOStateUpdateOne {
	if s != nil {
		ssuo.SetState(*s)
	}
	return ssuo
}

// SetData sets the "data" field.
func (ssuo *SSOStateUpdateOne) SetData(s string) *SSOStateUpdateOne {
	ssuo.mutation.SetData(s)
	return ssuo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (ssuo *SSOStateUpdateOne) SetNillableData(s *string) *SSOStateUpdateOne {
	if s != nil {
		ssuo.SetData(*s)
	}
	return ssuo
}

// SetExpiresAt sets the "expires_at" field.
func (ssuo *SSOStateUpdateOne) SetExpiresAt(t time.Time) *SSOStateUpdateOne {
	ssuo.mutation.SetExpiresAt(t)
	return ssuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (ssuo *SSOStateUpdateOne) SetNillableExpiresAt(t *time.Time) *SSOStateUpdateOne {
	if t != nil {
		ssuo.SetExpiresAt(*t)
	}
	return ssuo
}

// Mutation returns the SSOStateMutation object of the builder.
func (ssuo *SSOStateUpdateOne) Mutation() *SSOStateMutation {
	return ssuo.mutation
}

// Where appends a list predicates to the SSOStateUpdate builder.
func (ssuo *SSOStateUpdateOne) Where(ps ...predicate.SSOState) *SSOStateUpdateOne {
	ssuo.mutation.Where(ps...)
	return ssuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ssuo *SSOStateUpdateOne) Select(field string, fields ...string) *SSOStateUpdateOne {
	ssuo.fields = append([]string{field}, fields...)
	return ssuo
}

// Save executes the query and returns the updated SSOState entity.
func (ssuo *SSOStateUpdateOne) Save(ctx context.Context) (*SSOState, error) {
	ssuo.defaults()
	return withHooks(ctx, ssuo.sqlSave, ssuo.mutation, ssuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssuo *SSOStateUpdateOne) SaveX(ctx context.Context) *SSOState {
	node, err := ssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ssuo *SSOStateUpdateOne) Exec(ctx context.Context) error {
	_, err := ssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssuo *SSOStateUpdateOne) ExecX(ctx context.Context) {
	if err := ssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ssuo *SSOStateUpdateOne) defaults() {
	if _, ok := ssuo.mutation.UpdatedAt(); !ok {
		v := ssostate.UpdateDefaultUpdatedAt()
		ssuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ssuo *SSOStateUpdateOne) check() error {
	if v, ok := ssuo.mutation.Data(); ok {
		if err := ssostate.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "SSOState.data": %w`, err)}
		}
	}
	return nil
}

func (ssuo *SSOStateUpdateOne) sqlSave(ctx context.Context) (_node *SSOState, err error) {
	if err := ssuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ssostate.Table, ssostate.Columns, sqlgraph.NewFieldSpec(ssostate.FieldID, field.TypeString))
	id, ok := ssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SSOState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ssostate.FieldID)
		for _, f := range fields {
			if !ssostate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ssostate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ssuo.mutation.UpdatedAt(); ok {
		_spec.SetField(ssostate.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ssuo.mutation.State(); ok {
		_spec.SetField(ssostate.FieldState, field.TypeString, value)
	}
	if value, ok := ssuo.mutation.Data(); ok {
		_spec.SetField(ssostate.FieldData, field.TypeString, value)
	}
	if value, ok := ssuo.mutation.ExpiresAt(); ok {
		_spec.SetField(ssostate.FieldExpiresAt, field.TypeTime, value)
	}
	_node = &SSOState{config: ssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ssostate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ssuo.mutation.done = true
	return _node, nil
}
