package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Settlement holds the schema definition for the Settlement entity.
type Settlement struct {
	ent.Schema
}

// Fields of the Settlement.
func (Settlement) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x"),
		field.Int("y"),
		field.Bool("is_city"),
	}
}

// Edges of the Settlement.
func (Settlement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hexes", Hex.Type),
		edge.To("harbor", Harbor.Type).Unique(),
	}
}
