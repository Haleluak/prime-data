package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CasbinRule holds the schema definition for the CasbinRule entity.
type CasbinRule struct {
	ent.Schema
}

// Fields of the CasbinRule.
func (CasbinRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("p_type"),
		field.String("v0"),
		field.String("v1"),
		field.String("v2"),
		field.String("v3"),
		field.String("v4"),
		field.String("v5"),
	}
}

// Edges of the CasbinRule.
func (CasbinRule) Edges() []ent.Edge {
	return nil
}
