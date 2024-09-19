package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Words holds the schema definition for the Words entity.
type Words struct {
	ent.Schema
}

// Fields of the Words.
func (Words) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("group"),
		field.String("word"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Words.
func (Words) Edges() []ent.Edge {
	return nil
}
