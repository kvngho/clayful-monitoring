// Code generated by ent, DO NOT EDIT.

package coupon

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Coupon {
	return predicate.Coupon(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Coupon {
	return predicate.Coupon(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Coupon {
	return predicate.Coupon(sql.FieldLTE(FieldID, id))
}

// ClayfulID applies equality check predicate on the "clayful_id" field. It's identical to ClayfulIDEQ.
func ClayfulID(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldClayfulID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldName, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldEndDate, v))
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldActive, v))
}

// ClayfulIDEQ applies the EQ predicate on the "clayful_id" field.
func ClayfulIDEQ(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldClayfulID, v))
}

// ClayfulIDNEQ applies the NEQ predicate on the "clayful_id" field.
func ClayfulIDNEQ(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldNEQ(FieldClayfulID, v))
}

// ClayfulIDIn applies the In predicate on the "clayful_id" field.
func ClayfulIDIn(vs ...string) predicate.Coupon {
	return predicate.Coupon(sql.FieldIn(FieldClayfulID, vs...))
}

// ClayfulIDNotIn applies the NotIn predicate on the "clayful_id" field.
func ClayfulIDNotIn(vs ...string) predicate.Coupon {
	return predicate.Coupon(sql.FieldNotIn(FieldClayfulID, vs...))
}

// ClayfulIDGT applies the GT predicate on the "clayful_id" field.
func ClayfulIDGT(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldGT(FieldClayfulID, v))
}

// ClayfulIDGTE applies the GTE predicate on the "clayful_id" field.
func ClayfulIDGTE(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldGTE(FieldClayfulID, v))
}

// ClayfulIDLT applies the LT predicate on the "clayful_id" field.
func ClayfulIDLT(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldLT(FieldClayfulID, v))
}

// ClayfulIDLTE applies the LTE predicate on the "clayful_id" field.
func ClayfulIDLTE(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldLTE(FieldClayfulID, v))
}

// ClayfulIDContains applies the Contains predicate on the "clayful_id" field.
func ClayfulIDContains(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldContains(FieldClayfulID, v))
}

// ClayfulIDHasPrefix applies the HasPrefix predicate on the "clayful_id" field.
func ClayfulIDHasPrefix(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldHasPrefix(FieldClayfulID, v))
}

// ClayfulIDHasSuffix applies the HasSuffix predicate on the "clayful_id" field.
func ClayfulIDHasSuffix(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldHasSuffix(FieldClayfulID, v))
}

// ClayfulIDEqualFold applies the EqualFold predicate on the "clayful_id" field.
func ClayfulIDEqualFold(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldEqualFold(FieldClayfulID, v))
}

// ClayfulIDContainsFold applies the ContainsFold predicate on the "clayful_id" field.
func ClayfulIDContainsFold(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldContainsFold(FieldClayfulID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Coupon {
	return predicate.Coupon(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Coupon {
	return predicate.Coupon(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Coupon {
	return predicate.Coupon(sql.FieldContainsFold(FieldName, v))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Coupon {
	return predicate.Coupon(sql.FieldLTE(FieldEndDate, v))
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Coupon {
	return predicate.Coupon(sql.FieldEQ(FieldActive, v))
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Coupon {
	return predicate.Coupon(sql.FieldNEQ(FieldActive, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Coupon) predicate.Coupon {
	return predicate.Coupon(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Coupon) predicate.Coupon {
	return predicate.Coupon(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Coupon) predicate.Coupon {
	return predicate.Coupon(sql.NotPredicates(p))
}
