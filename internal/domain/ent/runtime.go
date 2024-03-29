// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/product"
	"github.com/kvngho/clayful-monitoring/internal/domain/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[1].Descriptor()
	// product.DefaultName holds the default value on creation for the name field.
	product.DefaultName = productDescName.Default.(string)
	// productDescPriceDeeping is the schema descriptor for price_deeping field.
	productDescPriceDeeping := productFields[2].Descriptor()
	// product.PriceDeepingValidator is a validator for the "price_deeping" field. It is called by the builders before save.
	product.PriceDeepingValidator = productDescPriceDeeping.Validators[0].(func(int) error)
	// productDescClayfulID is the schema descriptor for clayful_id field.
	productDescClayfulID := productFields[3].Descriptor()
	// product.DefaultClayfulID holds the default value on creation for the clayful_id field.
	product.DefaultClayfulID = productDescClayfulID.Default.(string)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.IDValidator is a validator for the "id" field. It is called by the builders before save.
	product.IDValidator = productDescID.Validators[0].(func(int) error)
}
