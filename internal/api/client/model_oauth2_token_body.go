/*
 * Frank Auth Server API
 *
 * Frank is a comprehensive authentication and authorization server that provides OAuth2, passwordless login, MFA,  passkeys, SSO, enterprise SSO, webhooks, organizations, and API keys for machine-to-machine authentication.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Oauth2TokenBody struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code,omitempty"`
	RedirectUri  string `json:"redirect_uri,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	CodeVerifier string `json:"code_verifier,omitempty"`
	Scope        string `json:"scope,omitempty"`
}
