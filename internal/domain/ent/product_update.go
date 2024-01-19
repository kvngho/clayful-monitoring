// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/predicate"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/product"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableName(s *string) *ProductUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetClayfulID sets the "clayful_id" field.
func (pu *ProductUpdate) SetClayfulID(s string) *ProductUpdate {
	pu.mutation.SetClayfulID(s)
	return pu
}

// SetNillableClayfulID sets the "clayful_id" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableClayfulID(s *string) *ProductUpdate {
	if s != nil {
		pu.SetClayfulID(*s)
	}
	return pu
}

// SetClayfulOptions sets the "clayful_options" field.
func (pu *ProductUpdate) SetClayfulOptions(a any) *ProductUpdate {
	pu.mutation.SetClayfulOptions(a)
	return pu
}

// ClearClayfulOptions clears the value of the "clayful_options" field.
func (pu *ProductUpdate) ClearClayfulOptions() *ProductUpdate {
	pu.mutation.ClearClayfulOptions()
	return pu
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.ClayfulID(); ok {
		_spec.SetField(product.FieldClayfulID, field.TypeString, value)
	}
	if value, ok := pu.mutation.ClayfulOptions(); ok {
		_spec.SetField(product.FieldClayfulOptions, field.TypeJSON, value)
	}
	if pu.mutation.ClayfulOptionsCleared() {
		_spec.ClearField(product.FieldClayfulOptions, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableName(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetClayfulID sets the "clayful_id" field.
func (puo *ProductUpdateOne) SetClayfulID(s string) *ProductUpdateOne {
	puo.mutation.SetClayfulID(s)
	return puo
}

// SetNillableClayfulID sets the "clayful_id" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableClayfulID(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetClayfulID(*s)
	}
	return puo
}

// SetClayfulOptions sets the "clayful_options" field.
func (puo *ProductUpdateOne) SetClayfulOptions(a any) *ProductUpdateOne {
	puo.mutation.SetClayfulOptions(a)
	return puo
}

// ClearClayfulOptions clears the value of the "clayful_options" field.
func (puo *ProductUpdateOne) ClearClayfulOptions() *ProductUpdateOne {
	puo.mutation.ClearClayfulOptions()
	return puo
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (puo *ProductUpdateOne) Where(ps ...predicate.Product) *ProductUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.ClayfulID(); ok {
		_spec.SetField(product.FieldClayfulID, field.TypeString, value)
	}
	if value, ok := puo.mutation.ClayfulOptions(); ok {
		_spec.SetField(product.FieldClayfulOptions, field.TypeJSON, value)
	}
	if puo.mutation.ClayfulOptionsCleared() {
		_spec.ClearField(product.FieldClayfulOptions, field.TypeJSON)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
