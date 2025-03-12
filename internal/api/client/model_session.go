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

type Session struct {
	Id             string    `json:"id,omitempty"`
	UserId         string    `json:"user_id,omitempty"`
	IpAddress      string    `json:"ip_address,omitempty"`
	UserAgent      string    `json:"user_agent,omitempty"`
	DeviceId       string    `json:"device_id,omitempty"`
	Location       string    `json:"location,omitempty"`
	OrganizationId string    `json:"organization_id,omitempty"`
	ExpiresAt      time.Time `json:"expires_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	LastActiveAt   time.Time `json:"last_active_at,omitempty"`
	IsActive       bool      `json:"is_active,omitempty"`
}
