package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StructTag(`json:oid,omitempty`),
		field.String("title"),
		field.Bool("is_done").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

func (Todo) Edges() []ent.Edge {
	return nil
}
