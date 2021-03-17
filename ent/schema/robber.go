package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Robber holds the schema definition for the Robber entity.
type Robber struct {
	ent.Schema
}

// Fields of the Robber.
func (Robber) Fields() []ent.Field {
	return nil
}

// Edges of the Robber.
func (Robber) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hex", Hex.Type).Unique(),
	}
}
