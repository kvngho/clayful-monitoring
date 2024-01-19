// Code generated by ent, DO NOT EDIT.

package product

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// ClayfulID applies equality check predicate on the "clayful_id" field. It's identical to ClayfulIDEQ.
func ClayfulID(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldClayfulID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// ClayfulIDEQ applies the EQ predicate on the "clayful_id" field.
func ClayfulIDEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldClayfulID, v))
}

// ClayfulIDNEQ applies the NEQ predicate on the "clayful_id" field.
func ClayfulIDNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldClayfulID, v))
}

// ClayfulIDIn applies the In predicate on the "clayful_id" field.
func ClayfulIDIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldClayfulID, vs...))
}

// ClayfulIDNotIn applies the NotIn predicate on the "clayful_id" field.
func ClayfulIDNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldClayfulID, vs...))
}

// ClayfulIDGT applies the GT predicate on the "clayful_id" field.
func ClayfulIDGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldClayfulID, v))
}

// ClayfulIDGTE applies the GTE predicate on the "clayful_id" field.
func ClayfulIDGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldClayfulID, v))
}

// ClayfulIDLT applies the LT predicate on the "clayful_id" field.
func ClayfulIDLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldClayfulID, v))
}

// ClayfulIDLTE applies the LTE predicate on the "clayful_id" field.
func ClayfulIDLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldClayfulID, v))
}

// ClayfulIDContains applies the Contains predicate on the "clayful_id" field.
func ClayfulIDContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldClayfulID, v))
}

// ClayfulIDHasPrefix applies the HasPrefix predicate on the "clayful_id" field.
func ClayfulIDHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldClayfulID, v))
}

// ClayfulIDHasSuffix applies the HasSuffix predicate on the "clayful_id" field.
func ClayfulIDHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldClayfulID, v))
}

// ClayfulIDEqualFold applies the EqualFold predicate on the "clayful_id" field.
func ClayfulIDEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldClayfulID, v))
}

// ClayfulIDContainsFold applies the ContainsFold predicate on the "clayful_id" field.
func ClayfulIDContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldClayfulID, v))
}

// ClayfulOptionsIsNil applies the IsNil predicate on the "clayful_options" field.
func ClayfulOptionsIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldClayfulOptions))
}

// ClayfulOptionsNotNil applies the NotNil predicate on the "clayful_options" field.
func ClayfulOptionsNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldClayfulOptions))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
