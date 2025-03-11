package handlers

import (
	"net/http"

	"github.com/juicycleff/frank/config"
	"github.com/juicycleff/frank/internal/middleware"
	"github.com/juicycleff/frank/internal/organization"
	"github.com/juicycleff/frank/pkg/errors"
	"github.com/juicycleff/frank/pkg/logging"
	"github.com/juicycleff/frank/pkg/utils"
)

// OrganizationHandler handles organization operations
type OrganizationHandler struct {
	orgService organization.Service
	config     *config.Config
	logger     logging.Logger
}

// NewOrganizationHandler creates a new organization handler
func NewOrganizationHandler(
	orgService organization.Service,
	config *config.Config,
	logger logging.Logger,
) *OrganizationHandler {
	return &OrganizationHandler{
		orgService: orgService,
		config:     config,
		logger:     logger,
	}
}

// CreateOrganizationRequest represents the input for creating an organization
type CreateOrganizationRequest struct {
	Name      string                 `json:"name" validate:"required"`
	Slug      string                 `json:"slug,omitempty"`
	Domain    string                 `json:"domain,omitempty"`
	LogoURL   string                 `json:"logo_url,omitempty"`
	Plan      string                 `json:"plan,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	TrialDays int                    `json:"trial_days,omitempty"`
	Features  []string               `json:"features,omitempty"`
}

// UpdateOrganizationRequest represents the input for updating an organization
type UpdateOrganizationRequest struct {
	Name     *string                `json:"name,omitempty"`
	Domain   *string                `json:"domain,omitempty"`
	LogoURL  *string                `json:"logo_url,omitempty"`
	Plan     *string                `json:"plan,omitempty"`
	Active   *bool                  `json:"active,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// ListOrganizations handles listing organizations with pagination
func (h *OrganizationHandler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	offset := utils.ParseQueryInt(r, "offset", 0)
	limit := utils.ParseQueryInt(r, "limit", 20)
	search := r.URL.Query().Get("search")

	// Create list params
	params := organization.ListParams{
		Offset: offset,
		Limit:  limit,
		Search: search,
	}

	// List organizations
	organizations, total, err := h.orgService.List(r.Context(), params)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return response with pagination
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"data":  organizations,
		"total": total,
		"pagination": map[string]interface{}{
			"offset": offset,
			"limit":  limit,
			"total":  total,
		},
	})
}

// CreateOrganization handles creating a new organization
func (h *OrganizationHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context
	userID, ok := middleware.GetUserID(r)
	if !ok {
		utils.RespondError(w, errors.New(errors.CodeUnauthorized, "not authenticated"))
		return
	}

	// Parse input
	var input CreateOrganizationRequest
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Map to service input
	createInput := organization.CreateOrganizationInput{
		Name:      input.Name,
		Slug:      input.Slug,
		Domain:    input.Domain,
		LogoURL:   input.LogoURL,
		Plan:      input.Plan,
		OwnerID:   userID,
		Metadata:  input.Metadata,
		TrialDays: input.TrialDays,
		Features:  input.Features,
	}

	// Create organization
	newOrg, err := h.orgService.Create(r.Context(), createInput)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return created organization
	utils.RespondJSON(w, http.StatusCreated, newOrg)
}

// GetOrganization handles retrieving an organization by ID
func (h *OrganizationHandler) GetOrganization(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Get organization
	org, err := h.orgService.Get(r.Context(), orgID)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return organization
	utils.RespondJSON(w, http.StatusOK, org)
}

// UpdateOrganization handles updating an organization
func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Parse input
	var input UpdateOrganizationRequest
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Map to service input
	updateInput := organization.UpdateOrganizationInput{
		Name:     input.Name,
		Domain:   input.Domain,
		LogoURL:  input.LogoURL,
		Plan:     input.Plan,
		Active:   input.Active,
		Metadata: input.Metadata,
	}

	// Update organization
	updatedOrg, err := h.orgService.Update(r.Context(), orgID, updateInput)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return updated organization
	utils.RespondJSON(w, http.StatusOK, updatedOrg)
}

// DeleteOrganization handles deleting an organization
func (h *OrganizationHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Delete organization
	err := h.orgService.Delete(r.Context(), orgID)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return success
	utils.RespondJSON(w, http.StatusNoContent, nil)
}

// ListOrganizationMembers handles listing members of an organization
func (h *OrganizationHandler) ListOrganizationMembers(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Parse query parameters
	offset := utils.ParseQueryInt(r, "offset", 0)
	limit := utils.ParseQueryInt(r, "limit", 20)
	search := r.URL.Query().Get("search")

	// Create list params
	params := organization.ListParams{
		Offset: offset,
		Limit:  limit,
		Search: search,
	}

	// List members
	members, total, err := h.orgService.GetMembers(r.Context(), orgID, params)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return response with pagination
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"data":  members,
		"total": total,
		"pagination": map[string]interface{}{
			"offset": offset,
			"limit":  limit,
			"total":  total,
		},
	})
}

// AddOrganizationMemberRequest represents the input for adding a member to an organization
type AddOrganizationMemberRequest struct {
	UserID string   `json:"user_id" validate:"required"`
	Roles  []string `json:"roles" validate:"required"`
}

// AddOrganizationMember handles adding a user to an organization
func (h *OrganizationHandler) AddOrganizationMember(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Parse input
	var input AddOrganizationMemberRequest
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Add member
	err := h.orgService.AddMember(r.Context(), orgID, input.UserID, input.Roles)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return success
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Member added successfully",
	})
}

// RemoveOrganizationMember handles removing a user from an organization
func (h *OrganizationHandler) RemoveOrganizationMember(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Get user ID from path
	userID := utils.GetPathVar(r, "userId")
	if userID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "user ID is required"))
		return
	}

	// Remove member
	err := h.orgService.RemoveMember(r.Context(), orgID, userID)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return success
	utils.RespondJSON(w, http.StatusNoContent, nil)
}

// UpdateOrganizationMemberRequest represents the input for updating a member's roles
type UpdateOrganizationMemberRequest struct {
	Roles []string `json:"roles" validate:"required"`
}

// UpdateOrganizationMember handles updating a member's roles in an organization
func (h *OrganizationHandler) UpdateOrganizationMember(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Get user ID from path
	userID := utils.GetPathVar(r, "userId")
	if userID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "user ID is required"))
		return
	}

	// Parse input
	var input UpdateOrganizationMemberRequest
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Update member
	err := h.orgService.UpdateMember(r.Context(), orgID, userID, input.Roles)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return success
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Member updated successfully",
	})
}

// ListOrganizationFeatures handles listing features of an organization
func (h *OrganizationHandler) ListOrganizationFeatures(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Get features
	features, err := h.orgService.GetFeatures(r.Context(), orgID)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return features
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"data": features,
	})
}

// EnableOrganizationFeatureRequest represents the input for enabling a feature
type EnableOrganizationFeatureRequest struct {
	FeatureKey string                 `json:"feature_key" validate:"required"`
	Settings   map[string]interface{} `json:"settings,omitempty"`
}

// EnableOrganizationFeature handles enabling a feature for an organization
func (h *OrganizationHandler) EnableOrganizationFeature(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Parse input
	var input EnableOrganizationFeatureRequest
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Enable feature
	err := h.orgService.EnableFeature(r.Context(), orgID, input.FeatureKey, input.Settings)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return success
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Feature enabled successfully",
	})
}

// DisableOrganizationFeature handles disabling a feature for an organization
func (h *OrganizationHandler) DisableOrganizationFeature(w http.ResponseWriter, r *http.Request) {
	// Get organization ID from path
	orgID := utils.GetPathVar(r, "id")
	if orgID == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "organization ID is required"))
		return
	}

	// Get feature key from path
	featureKey := utils.GetPathVar(r, "featureKey")
	if featureKey == "" {
		utils.RespondError(w, errors.New(errors.CodeInvalidInput, "feature key is required"))
		return
	}

	// Disable feature
	err := h.orgService.DisableFeature(r.Context(), orgID, featureKey)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Return success
	utils.RespondJSON(w, http.StatusNoContent, nil)
}

// SetupRoutes sets up the organization routes
func (h *OrganizationHandler) SetupRoutes(router *http.ServeMux) {
	router.HandleFunc("/api/v1/organizations", h.ListOrganizations)
	router.HandleFunc("/api/v1/organizations", h.CreateOrganization)
	router.HandleFunc("/api/v1/organizations/{id}", h.GetOrganization)
	router.HandleFunc("/api/v1/organizations/{id}", h.UpdateOrganization)
	router.HandleFunc("/api/v1/organizations/{id}", h.DeleteOrganization)
	router.HandleFunc("/api/v1/organizations/{id}/members", h.ListOrganizationMembers)
	router.HandleFunc("/api/v1/organizations/{id}/members", h.AddOrganizationMember)
	router.HandleFunc("/api/v1/organizations/{id}/members/{userId}", h.RemoveOrganizationMember)
	router.HandleFunc("/api/v1/organizations/{id}/members/{userId}", h.UpdateOrganizationMember)
	router.HandleFunc("/api/v1/organizations/{id}/features", h.ListOrganizationFeatures)
	router.HandleFunc("/api/v1/organizations/{id}/features", h.EnableOrganizationFeature)
	router.HandleFunc("/api/v1/organizations/{id}/features/{featureKey}", h.DisableOrganizationFeature)
}

// Static handler functions for direct router registration

// ListOrganizations handles listing organizations API endpoint
func ListOrganizations(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.ListOrganizations(w, r)
}

// CreateOrganization handles creating an organization API endpoint
func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.CreateOrganization(w, r)
}

// GetOrganization handles retrieving an organization API endpoint
func GetOrganization(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.GetOrganization(w, r)
}

// UpdateOrganization handles updating an organization API endpoint
func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.UpdateOrganization(w, r)
}

// DeleteOrganization handles deleting an organization API endpoint
func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.DeleteOrganization(w, r)
}

// ListOrganizationMembers handles listing organization members API endpoint
func ListOrganizationMembers(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.ListOrganizationMembers(w, r)
}

// AddOrganizationMember handles adding an organization member API endpoint
func AddOrganizationMember(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.AddOrganizationMember(w, r)
}

// RemoveOrganizationMember handles removing an organization member API endpoint
func RemoveOrganizationMember(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.RemoveOrganizationMember(w, r)
}

// UpdateOrganizationMember handles updating an organization member API endpoint
func UpdateOrganizationMember(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.UpdateOrganizationMember(w, r)
}

// ListOrganizationFeatures handles listing organization features API endpoint
func ListOrganizationFeatures(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.ListOrganizationFeatures(w, r)
}

// EnableOrganizationFeature handles enabling an organization feature API endpoint
func EnableOrganizationFeature(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.EnableOrganizationFeature(w, r)
}

// DisableOrganizationFeature handles disabling an organization feature API endpoint
func DisableOrganizationFeature(w http.ResponseWriter, r *http.Request) {
	HandlerFromContext(r.Context()).Organization.DisableOrganizationFeature(w, r)
}
