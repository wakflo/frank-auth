/*
 * Frank Auth Server API
 *
 * Frank is a comprehensive authentication and authorization server that provides OAuth2, passwordless login, MFA,  passkeys, SSO, enterprise SSO, webhooks, organizations, and API keys for machine-to-machine authentication.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrganizationsBody struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug,omitempty"`
	Domain    string   `json:"domain,omitempty"`
	LogoUrl   string   `json:"logo_url,omitempty"`
	Plan      string   `json:"plan,omitempty"`
	Metadata  ModelMap `json:"metadata,omitempty"`
	TrialDays int32    `json:"trial_days,omitempty"`
	Features  []string `json:"features,omitempty"`
}
