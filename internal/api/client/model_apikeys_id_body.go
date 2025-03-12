/*
 * Frank Auth Server API
 *
 * Frank is a comprehensive authentication and authorization server that provides OAuth2, passwordless login, MFA,  passkeys, SSO, enterprise SSO, webhooks, organizations, and API keys for machine-to-machine authentication.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

type ApikeysIdBody struct {
	Name        string    `json:"name,omitempty"`
	Active      bool      `json:"active,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	Scopes      []string  `json:"scopes,omitempty"`
	Metadata    ModelMap  `json:"metadata,omitempty"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
}
