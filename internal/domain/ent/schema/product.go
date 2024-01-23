package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "product_product"},
    }
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name").Default(""),
		field.Int("price_deeping").Positive(),
		field.String("clayful_id").Default(""),
		field.Any("clayful_options").Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}

func (Product) Table() string {
    return "product_product"
}
