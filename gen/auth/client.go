// Code generated by goa v3.20.0, DO NOT EDIT.
//
// auth client
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package auth

import (
	"context"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	goa "goa.design/goa/v3/pkg"
)

// Client is the "auth" service client.
type Client struct {
	LoginEndpoint          goa.Endpoint
	RegisterEndpoint       goa.Endpoint
	LogoutEndpoint         goa.Endpoint
	RefreshTokenEndpoint   goa.Endpoint
	ForgotPasswordEndpoint goa.Endpoint
	ResetPasswordEndpoint  goa.Endpoint
	VerifyEmailEndpoint    goa.Endpoint
	MeEndpoint             goa.Endpoint
	CsrfEndpoint           goa.Endpoint
}

// NewClient initializes a "auth" service client given the endpoints.
func NewClient(login, register, logout, refreshToken, forgotPassword, resetPassword, verifyEmail, me, csrf goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint:          login,
		RegisterEndpoint:       register,
		LogoutEndpoint:         logout,
		RefreshTokenEndpoint:   refreshToken,
		ForgotPasswordEndpoint: forgotPassword,
		ResetPasswordEndpoint:  resetPassword,
		VerifyEmailEndpoint:    verifyEmail,
		MeEndpoint:             me,
		CsrfEndpoint:           csrf,
	}
}

// Login calls the "login" endpoint of the "auth" service.
// Login may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "forbidden" (type *ForbiddenError): Account is locked or email not verified
//   - "unauthorized" (type *UnauthorizedError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginResponse, err error) {
	var ires any
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResponse), nil
}

// Register calls the "register" endpoint of the "auth" service.
// Register may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "conflict" (type *ConflictError): Email already exists
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res *LoginResponse, err error) {
	var ires any
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResponse), nil
}

// Logout calls the "logout" endpoint of the "auth" service.
// Logout may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) Logout(ctx context.Context, p *LogoutPayload) (res *LogoutResult, err error) {
	var ires any
	ires, err = c.LogoutEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LogoutResult), nil
}

// RefreshToken calls the "refresh_token" endpoint of the "auth" service.
// RefreshToken may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError): Invalid refresh token
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) RefreshToken(ctx context.Context, p *RefreshTokenPayload) (res *RefreshTokenResponse, err error) {
	var ires any
	ires, err = c.RefreshTokenEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*RefreshTokenResponse), nil
}

// ForgotPassword calls the "forgot_password" endpoint of the "auth" service.
// ForgotPassword may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) ForgotPassword(ctx context.Context, p *ForgotPasswordPayload) (res *ForgotPasswordResult, err error) {
	var ires any
	ires, err = c.ForgotPasswordEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ForgotPasswordResult), nil
}

// ResetPassword calls the "reset_password" endpoint of the "auth" service.
// ResetPassword may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError): Invalid or expired token
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) ResetPassword(ctx context.Context, p *ResetPasswordPayload) (res *ResetPasswordResult, err error) {
	var ires any
	ires, err = c.ResetPasswordEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResetPasswordResult), nil
}

// VerifyEmail calls the "verify_email" endpoint of the "auth" service.
// VerifyEmail may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) VerifyEmail(ctx context.Context, p *VerifyEmailPayload) (res *VerifyEmailResult, err error) {
	var ires any
	ires, err = c.VerifyEmailEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*VerifyEmailResult), nil
}

// Me calls the "me" endpoint of the "auth" service.
// Me may return the following errors:
//   - "unauthorized" (type *UnauthorizedError)
//   - "bad_request" (type *BadRequestError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) Me(ctx context.Context, p *MePayload) (res *designtypes.User, err error) {
	var ires any
	ires, err = c.MeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*designtypes.User), nil
}

// Csrf calls the "csrf" endpoint of the "auth" service.
// Csrf may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - "invalid_credentials" (type *InternalServerError)
//   - error: internal error
func (c *Client) Csrf(ctx context.Context, p *CsrfPayload) (res *CSRFTokenResponse, err error) {
	var ires any
	ires, err = c.CsrfEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CSRFTokenResponse), nil
}
