package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Verification holds the schema definition for the Verification entity.
type Verification struct {
	ent.Schema
}

// Fields of the Verification.
func (Verification) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique(),
		field.String("user_id").
			NotEmpty(),
		field.String("type").
			NotEmpty().
			Comment("Verification type: email, phone, password_reset, magic_link"),
		field.String("token").
			Unique().
			NotEmpty().
			Sensitive(),
		field.String("email").
			Optional(),
		field.String("phone_number").
			Optional(),
		field.String("redirect_url").
			Optional(),
		field.Bool("used").
			Default(false),
		field.Time("used_at").
			Optional().
			Nillable(),
		field.Int("attempts").
			Default(0),
		field.Time("expires_at"),
		field.String("ip_address").
			Optional(),
		field.String("user_agent").
			Optional(),
		field.JSON("metadata", map[string]interface{}{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Verification.
func (Verification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("verifications").
			Field("user_id").
			Unique().
			Required(),
	}
}

// Indexes of the Verification.
func (Verification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("token"),
		index.Fields("email"),
		index.Fields("phone_number"),
		index.Fields("expires_at"),
	}
}
