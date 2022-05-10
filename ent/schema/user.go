package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/segmentio/ksuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			return ksuid.New().String()
		}),
		field.String("name").Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
