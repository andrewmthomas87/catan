package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Harbor holds the schema definition for the Harbor entity.
type Harbor struct {
	ent.Schema
}

// Fields of the Harbor.
func (Harbor) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x"),
		field.Int("y"),
		field.Enum("resource").Values("generic", "brick", "lumber", "ore", "grain", "wool"),
	}
}

// Edges of the Harbor.
func (Harbor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("settlement", Settlement.Type).Ref("harbor").Unique(),
	}
}
