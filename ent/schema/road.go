package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Road holds the schema definition for the Road entity.
type Road struct {
	ent.Schema
}

// Fields of the Road.
func (Road) Fields() []ent.Field {
	return []ent.Field{
		field.Int("x"),
		field.Int("y"),
	}
}

// Edges of the Road.
func (Road) Edges() []ent.Edge {
	return nil
}
