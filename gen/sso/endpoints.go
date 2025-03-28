// Code generated by goa v3.20.0, DO NOT EDIT.
//
// sso endpoints
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package sso

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "sso" service endpoints.
type Endpoints struct {
	ListProviders          goa.Endpoint
	ProviderAuth           goa.Endpoint
	ProviderCallback       goa.Endpoint
	ListIdentityProviders  goa.Endpoint
	CreateIdentityProvider goa.Endpoint
	GetIdentityProvider    goa.Endpoint
	UpdateIdentityProvider goa.Endpoint
	DeleteIdentityProvider goa.Endpoint
	SamlMetadata           goa.Endpoint
	SamlAcs                goa.Endpoint
}

// NewEndpoints wraps the methods of the "sso" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ListProviders:          NewListProvidersEndpoint(s, a.OAuth2Auth, a.APIKeyAuth, a.JWTAuth),
		ProviderAuth:           NewProviderAuthEndpoint(s, a.OAuth2Auth, a.APIKeyAuth, a.JWTAuth),
		ProviderCallback:       NewProviderCallbackEndpoint(s, a.OAuth2Auth, a.APIKeyAuth, a.JWTAuth),
		ListIdentityProviders:  NewListIdentityProvidersEndpoint(s, a.JWTAuth),
		CreateIdentityProvider: NewCreateIdentityProviderEndpoint(s, a.JWTAuth),
		GetIdentityProvider:    NewGetIdentityProviderEndpoint(s, a.JWTAuth),
		UpdateIdentityProvider: NewUpdateIdentityProviderEndpoint(s, a.JWTAuth),
		DeleteIdentityProvider: NewDeleteIdentityProviderEndpoint(s, a.JWTAuth),
		SamlMetadata:           NewSamlMetadataEndpoint(s),
		SamlAcs:                NewSamlAcsEndpoint(s),
	}
}

// Use applies the given middleware to all the "sso" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListProviders = m(e.ListProviders)
	e.ProviderAuth = m(e.ProviderAuth)
	e.ProviderCallback = m(e.ProviderCallback)
	e.ListIdentityProviders = m(e.ListIdentityProviders)
	e.CreateIdentityProvider = m(e.CreateIdentityProvider)
	e.GetIdentityProvider = m(e.GetIdentityProvider)
	e.UpdateIdentityProvider = m(e.UpdateIdentityProvider)
	e.DeleteIdentityProvider = m(e.DeleteIdentityProvider)
	e.SamlMetadata = m(e.SamlMetadata)
	e.SamlAcs = m(e.SamlAcs)
}

// NewListProvidersEndpoint returns an endpoint function that calls the method
// "list_providers" of service "sso".
func NewListProvidersEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authAPIKeyFn security.AuthAPIKeyFunc, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListProvidersPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"profile", "email", "openid", "offline_access", "api"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:             "authorization_code",
					AuthorizationURL: "/v1/oauth/authorize",
					TokenURL:         "/v1/oauth/token",
					RefreshURL:       "/v1/oauth/refresh",
				},
			},
		}
		var token string
		if p.Oauth2 != nil {
			token = *p.Oauth2
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.XAPIKey != nil {
				key = *p.XAPIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWT != nil {
				token = *p.JWT
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.ListProviders(ctx, p)
	}
}

// NewProviderAuthEndpoint returns an endpoint function that calls the method
// "provider_auth" of service "sso".
func NewProviderAuthEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authAPIKeyFn security.AuthAPIKeyFunc, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ProviderAuthPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"profile", "email", "openid", "offline_access", "api"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:             "authorization_code",
					AuthorizationURL: "/v1/oauth/authorize",
					TokenURL:         "/v1/oauth/token",
					RefreshURL:       "/v1/oauth/refresh",
				},
			},
		}
		var token string
		if p.Oauth2 != nil {
			token = *p.Oauth2
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.XAPIKey != nil {
				key = *p.XAPIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWT != nil {
				token = *p.JWT
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return nil, s.ProviderAuth(ctx, p)
	}
}

// NewProviderCallbackEndpoint returns an endpoint function that calls the
// method "provider_callback" of service "sso".
func NewProviderCallbackEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authAPIKeyFn security.AuthAPIKeyFunc, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ProviderCallbackPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"profile", "email", "openid", "offline_access", "api"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:             "authorization_code",
					AuthorizationURL: "/v1/oauth/authorize",
					TokenURL:         "/v1/oauth/token",
					RefreshURL:       "/v1/oauth/refresh",
				},
			},
		}
		var token string
		if p.Oauth2 != nil {
			token = *p.Oauth2
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{},
			}
			var key string
			if p.XAPIKey != nil {
				key = *p.XAPIKey
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWT != nil {
				token = *p.JWT
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.ProviderCallback(ctx, p)
	}
}

// NewListIdentityProvidersEndpoint returns an endpoint function that calls the
// method "list_identity_providers" of service "sso".
func NewListIdentityProvidersEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListIdentityProvidersPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ListIdentityProviders(ctx, p)
	}
}

// NewCreateIdentityProviderEndpoint returns an endpoint function that calls
// the method "create_identity_provider" of service "sso".
func NewCreateIdentityProviderEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateIdentityProviderPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.CreateIdentityProvider(ctx, p)
	}
}

// NewGetIdentityProviderEndpoint returns an endpoint function that calls the
// method "get_identity_provider" of service "sso".
func NewGetIdentityProviderEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetIdentityProviderPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetIdentityProvider(ctx, p)
	}
}

// NewUpdateIdentityProviderEndpoint returns an endpoint function that calls
// the method "update_identity_provider" of service "sso".
func NewUpdateIdentityProviderEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateIdentityProviderPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.UpdateIdentityProvider(ctx, p)
	}
}

// NewDeleteIdentityProviderEndpoint returns an endpoint function that calls
// the method "delete_identity_provider" of service "sso".
func NewDeleteIdentityProviderEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteIdentityProviderPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteIdentityProvider(ctx, p)
	}
}

// NewSamlMetadataEndpoint returns an endpoint function that calls the method
// "saml_metadata" of service "sso".
func NewSamlMetadataEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SamlMetadataPayload)
		return s.SamlMetadata(ctx, p)
	}
}

// NewSamlAcsEndpoint returns an endpoint function that calls the method
// "saml_acs" of service "sso".
func NewSamlAcsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SamlAcsPayload)
		return s.SamlAcs(ctx, p)
	}
}
