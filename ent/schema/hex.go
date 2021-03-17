package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hex holds the schema definition for the Hex entity.
type Hex struct {
	ent.Schema
}

// Fields of the Hex.
func (Hex) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x"),
		field.Int("y"),
		field.Enum("terrain").Values("hills", "forest", "mountains", "fields", "pasture", "desert"),
	}
}

// Edges of the Hex.
func (Hex) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("number_token", NumberToken.Type).Ref("hex").Unique(),
		edge.From("robber", Robber.Type).Ref("hex").Unique(),
		edge.From("settlements", Settlement.Type).Ref("hexes"),
	}
}
