package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/juicycleff/frank/pkg/errors"
	"github.com/juicycleff/frank/pkg/utils"
)

// Handlers contains all handler instances
type Handlers struct {
	Auth         *AuthHandler
	User         *UserHandler
	Organization *OrganizationHandler
	OAuth        *OAuthHandler
	Webhook      *WebhookHandler
	APIKey       *APIKeyHandler
	MFA          *MFAHandler
	Passkey      *PasskeyHandler
	SSO          *SSOHandler
	Passwordless *PasswordlessHandler
}

// ContextKey is a type for context keys
type ContextKey string

// HandlerContextKey is the key for handlers in context
const HandlerContextKey ContextKey = "handlers"

// SetHandlersContext adds handlers to context
func SetHandlersContext(ctx context.Context, handlers *Handlers) context.Context {
	return context.WithValue(ctx, HandlerContextKey, handlers)
}

// HandlerFromContext gets handlers from context
func HandlerFromContext(ctx context.Context) *Handlers {
	if handlers, ok := ctx.Value(HandlerContextKey).(*Handlers); ok {
		return handlers
	}
	return nil
}

// HandlerContextMiddleware adds handlers to request context
func HandlerContextMiddleware(handlers *Handlers) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := SetHandlersContext(r.Context(), handlers)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// HealthCheck handles the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"status": "healthy",
	})
}

// ReadyCheck handles the readiness check endpoint
func ReadyCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"status": "ready",
	})
}

// NotFound handles 404 errors
func NotFound(w http.ResponseWriter, r *http.Request) {
	utils.RespondError(w, errors.New(errors.CodeNotFound, "resource not found"))
}

// MethodNotAllowed handles 405 errors
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	utils.RespondError(w, errors.New(errors.CodeMethodNotAllowed, fmt.Sprintf("method %s not allowed", r.Method)))
}

// SAMLAssertionConsumerService handles SAML assertion consumer service endpoint
func SAMLAssertionConsumerService(w http.ResponseWriter, r *http.Request) {
	// Delegate to SSO service if available
	if handlers := HandlerFromContext(r.Context()); handlers != nil && handlers.SSO != nil {
		// The actual implementation would need to handle SAML-specific logic
		// For now, just return an error
		utils.RespondError(w, errors.New(errors.CodeNotImplemented, "SAML ACS endpoint not implemented"))
		return
	}

	utils.RespondError(w, errors.New(errors.CodeNotImplemented, "SAML ACS endpoint not implemented"))
}

// SAMLMetadata handles SAML metadata endpoint
func SAMLMetadata(w http.ResponseWriter, r *http.Request) {
	// Delegate to SSO service if available
	if handlers := HandlerFromContext(r.Context()); handlers != nil && handlers.SSO != nil {
		// The actual implementation would need to handle SAML-specific logic
		// For now, just return an error
		utils.RespondError(w, errors.New(errors.CodeNotImplemented, "SAML metadata endpoint not implemented"))
		return
	}

	utils.RespondError(w, errors.New(errors.CodeNotImplemented, "SAML metadata endpoint not implemented"))
}
