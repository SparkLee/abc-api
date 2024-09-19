package schema

import "entgo.io/ent"

// Words holds the schema definition for the Words entity.
type Words struct {
	ent.Schema
}

// Fields of the Words.
func (Words) Fields() []ent.Field {
	return nil
}

// Edges of the Words.
func (Words) Edges() []ent.Edge {
	return nil
}
