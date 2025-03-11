package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OAuthToken holds the schema definition for the OAuthToken entity.
type OAuthToken struct {
	ent.Schema
}

// Fields of the OAuthToken.
func (OAuthToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique(),
		field.String("access_token").
			Unique().
			NotEmpty().
			Sensitive(),
		field.String("refresh_token").
			Unique().
			Optional().
			Sensitive(),
		field.String("token_type").
			Default("bearer"),
		field.String("client_id").
			NotEmpty(),
		field.String("user_id").
			NotEmpty(),
		field.String("organization_id").
			Optional(),
		field.Strings("scope_names").
			Optional().
			Comment("Scope names as strings for quick access"),
		field.Int("expires_in").
			Default(3600),
		field.Time("expires_at"),
		field.Time("refresh_token_expires_at").
			Optional().
			Nillable(),
		field.Bool("revoked").
			Default(false),
		field.Time("revoked_at").
			Optional().
			Nillable(),
		field.String("ip_address").
			Optional(),
		field.String("user_agent").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the OAuthToken.
func (OAuthToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client", OAuthClient.Type).
			Ref("tokens").
			Unique().
			Field("client_id").
			Required(),
		edge.From("user", User.Type).
			Ref("oauth_tokens").
			Unique().
			Field("user_id").
			Required(),
		edge.To("scopes", OAuthScope.Type),
	}
}

// Indexes of the OAuthToken.
func (OAuthToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("access_token"),
		index.Fields("refresh_token"),
		index.Fields("client_id"),
		index.Fields("user_id"),
		index.Fields("organization_id"),
		index.Fields("expires_at"),
	}
}
