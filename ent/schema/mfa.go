package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MFA holds the schema definition for the MFA entity.
type MFA struct {
	ent.Schema
}

// Fields of the MFA.
func (MFA) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique(),
		field.String("user_id").
			NotEmpty(),
		field.String("method").
			NotEmpty().
			Comment("The MFA method type: 'totp', 'sms', 'email', 'backup_codes'"),
		field.String("secret").
			NotEmpty().
			Sensitive(),
		field.Bool("verified").
			Default(false),
		field.Bool("active").
			Default(true),
		field.JSON("backup_codes", []string{}).
			Optional(),
		field.String("phone_number").
			Optional(),
		field.String("email").
			Optional(),
		field.Time("last_used").
			Optional().
			Nillable(),
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

// Edges of the MFA.
func (MFA) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("mfa_methods").
			Field("user_id").
			Unique().
			Required(),
	}
}

// Indexes of the MFA.
func (MFA) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("method", "user_id").
			Unique(),
	}
}
