// Code generated by ent, DO NOT EDIT.

package featureflag

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the featureflag type in the database.
	Label = "feature_flag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// FieldIsPremium holds the string denoting the is_premium field in the database.
	FieldIsPremium = "is_premium"
	// FieldComponent holds the string denoting the component field in the database.
	FieldComponent = "component"
	// EdgeOrganizations holds the string denoting the organizations edge name in mutations.
	EdgeOrganizations = "organizations"
	// Table holds the table name of the featureflag in the database.
	Table = "feature_flags"
	// OrganizationsTable is the table that holds the organizations relation/edge.
	OrganizationsTable = "organization_features"
	// OrganizationsInverseTable is the table name for the OrganizationFeature entity.
	// It exists in this package in order to avoid circular dependency with the "organizationfeature" package.
	OrganizationsInverseTable = "organization_features"
	// OrganizationsColumn is the table column denoting the organizations relation/edge.
	OrganizationsColumn = "feature_id"
)

// Columns holds all SQL columns for featureflag fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldKey,
	FieldDescription,
	FieldEnabled,
	FieldIsPremium,
	FieldComponent,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
	// DefaultIsPremium holds the default value on creation for the "is_premium" field.
	DefaultIsPremium bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// Component defines the type for the "component" enum field.
type Component string

// Component values.
const (
	ComponentOauth2       Component = "oauth2"
	ComponentPasswordless Component = "passwordless"
	ComponentMfa          Component = "mfa"
	ComponentPasskeys     Component = "passkeys"
	ComponentSSO          Component = "sso"
	ComponentEnterprise   Component = "enterprise"
	ComponentWebhooks     Component = "webhooks"
	ComponentAPIKeys      Component = "api_keys"
)

func (c Component) String() string {
	return string(c)
}

// ComponentValidator is a validator for the "component" field enum values. It is called by the builders before save.
func ComponentValidator(c Component) error {
	switch c {
	case ComponentOauth2, ComponentPasswordless, ComponentMfa, ComponentPasskeys, ComponentSSO, ComponentEnterprise, ComponentWebhooks, ComponentAPIKeys:
		return nil
	default:
		return fmt.Errorf("featureflag: invalid enum value for component field: %q", c)
	}
}

// OrderOption defines the ordering options for the FeatureFlag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByEnabled orders the results by the enabled field.
func ByEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnabled, opts...).ToFunc()
}

// ByIsPremium orders the results by the is_premium field.
func ByIsPremium(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPremium, opts...).ToFunc()
}

// ByComponent orders the results by the component field.
func ByComponent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComponent, opts...).ToFunc()
}

// ByOrganizationsCount orders the results by organizations count.
func ByOrganizationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrganizationsStep(), opts...)
	}
}

// ByOrganizations orders the results by organizations terms.
func ByOrganizations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrganizationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OrganizationsTable, OrganizationsColumn),
	)
}
