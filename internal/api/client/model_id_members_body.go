/*
 * Frank Auth Server API
 *
 * Frank is a comprehensive authentication and authorization server that provides OAuth2, passwordless login, MFA,  passkeys, SSO, enterprise SSO, webhooks, organizations, and API keys for machine-to-machine authentication.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type IdMembersBody struct {
	UserId string   `json:"user_id"`
	Roles  []string `json:"roles"`
}
