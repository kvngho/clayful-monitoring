package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Coupon holds the schema definition for the Coupon entity.
type Coupon struct {
	ent.Schema
}

func (Coupon) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "shop_couponall"},
    }
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("clayful_id").StorageKey("_id"),
		field.String("name"),
		field.Time("end_date"),
		field.Bool("active"),
	}
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return nil
}
