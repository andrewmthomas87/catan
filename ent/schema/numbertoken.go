package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NumberToken holds the schema definition for the NumberToken entity.
type NumberToken struct {
	ent.Schema
}

// Fields of the NumberToken.
func (NumberToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value").Min(1).Max(12),
	}
}

// Edges of the NumberToken.
func (NumberToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hex", Hex.Type).Unique(),
	}
}
