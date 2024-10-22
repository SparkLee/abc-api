package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Words holds the schema definition for the Words entity.
type Words struct {
	ent.Schema
}

func (Words) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("单词表"),
	}
}

// Fields of the Words.
func (Words) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("group").Comment("分组"),
		field.String("word").Unique(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the Words.
func (Words) Edges() []ent.Edge {
	return nil
}
