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
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/coupon"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/predicate"
)

// CouponUpdate is the builder for updating Coupon entities.
type CouponUpdate struct {
	config
	hooks    []Hook
	mutation *CouponMutation
}

// Where appends a list predicates to the CouponUpdate builder.
func (cu *CouponUpdate) Where(ps ...predicate.Coupon) *CouponUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetClayfulID sets the "clayful_id" field.
func (cu *CouponUpdate) SetClayfulID(s string) *CouponUpdate {
	cu.mutation.SetClayfulID(s)
	return cu
}

// SetNillableClayfulID sets the "clayful_id" field if the given value is not nil.
func (cu *CouponUpdate) SetNillableClayfulID(s *string) *CouponUpdate {
	if s != nil {
		cu.SetClayfulID(*s)
	}
	return cu
}

// SetName sets the "name" field.
func (cu *CouponUpdate) SetName(s string) *CouponUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CouponUpdate) SetNillableName(s *string) *CouponUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetEndDate sets the "end_date" field.
func (cu *CouponUpdate) SetEndDate(t time.Time) *CouponUpdate {
	cu.mutation.SetEndDate(t)
	return cu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cu *CouponUpdate) SetNillableEndDate(t *time.Time) *CouponUpdate {
	if t != nil {
		cu.SetEndDate(*t)
	}
	return cu
}

// SetActive sets the "active" field.
func (cu *CouponUpdate) SetActive(b bool) *CouponUpdate {
	cu.mutation.SetActive(b)
	return cu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cu *CouponUpdate) SetNillableActive(b *bool) *CouponUpdate {
	if b != nil {
		cu.SetActive(*b)
	}
	return cu
}

// Mutation returns the CouponMutation object of the builder.
func (cu *CouponUpdate) Mutation() *CouponMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CouponUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CouponUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CouponUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CouponUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CouponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(coupon.Table, coupon.Columns, sqlgraph.NewFieldSpec(coupon.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.ClayfulID(); ok {
		_spec.SetField(coupon.FieldClayfulID, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(coupon.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.EndDate(); ok {
		_spec.SetField(coupon.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Active(); ok {
		_spec.SetField(coupon.FieldActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CouponUpdateOne is the builder for updating a single Coupon entity.
type CouponUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CouponMutation
}

// SetClayfulID sets the "clayful_id" field.
func (cuo *CouponUpdateOne) SetClayfulID(s string) *CouponUpdateOne {
	cuo.mutation.SetClayfulID(s)
	return cuo
}

// SetNillableClayfulID sets the "clayful_id" field if the given value is not nil.
func (cuo *CouponUpdateOne) SetNillableClayfulID(s *string) *CouponUpdateOne {
	if s != nil {
		cuo.SetClayfulID(*s)
	}
	return cuo
}

// SetName sets the "name" field.
func (cuo *CouponUpdateOne) SetName(s string) *CouponUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CouponUpdateOne) SetNillableName(s *string) *CouponUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetEndDate sets the "end_date" field.
func (cuo *CouponUpdateOne) SetEndDate(t time.Time) *CouponUpdateOne {
	cuo.mutation.SetEndDate(t)
	return cuo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cuo *CouponUpdateOne) SetNillableEndDate(t *time.Time) *CouponUpdateOne {
	if t != nil {
		cuo.SetEndDate(*t)
	}
	return cuo
}

// SetActive sets the "active" field.
func (cuo *CouponUpdateOne) SetActive(b bool) *CouponUpdateOne {
	cuo.mutation.SetActive(b)
	return cuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cuo *CouponUpdateOne) SetNillableActive(b *bool) *CouponUpdateOne {
	if b != nil {
		cuo.SetActive(*b)
	}
	return cuo
}

// Mutation returns the CouponMutation object of the builder.
func (cuo *CouponUpdateOne) Mutation() *CouponMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CouponUpdate builder.
func (cuo *CouponUpdateOne) Where(ps ...predicate.Coupon) *CouponUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CouponUpdateOne) Select(field string, fields ...string) *CouponUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Coupon entity.
func (cuo *CouponUpdateOne) Save(ctx context.Context) (*Coupon, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CouponUpdateOne) SaveX(ctx context.Context) *Coupon {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CouponUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CouponUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CouponUpdateOne) sqlSave(ctx context.Context) (_node *Coupon, err error) {
	_spec := sqlgraph.NewUpdateSpec(coupon.Table, coupon.Columns, sqlgraph.NewFieldSpec(coupon.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Coupon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coupon.FieldID)
		for _, f := range fields {
			if !coupon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.ClayfulID(); ok {
		_spec.SetField(coupon.FieldClayfulID, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(coupon.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.EndDate(); ok {
		_spec.SetField(coupon.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Active(); ok {
		_spec.SetField(coupon.FieldActive, field.TypeBool, value)
	}
	_node = &Coupon{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
