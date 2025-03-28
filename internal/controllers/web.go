package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/juicycleff/frank/config"
	webhttp "github.com/juicycleff/frank/gen/http/web/server"
	websvc "github.com/juicycleff/frank/gen/web"
	"github.com/juicycleff/frank/internal/middleware"
	"github.com/juicycleff/frank/internal/services"
	"github.com/juicycleff/frank/pkg/astro_fs"
	"github.com/juicycleff/frank/pkg/data"
	"github.com/juicycleff/frank/pkg/logging"
	"github.com/juicycleff/frank/pkg/utils"
	goahttp "goa.design/goa/v3/http"
)

type WebService struct {
	clients *data.Clients
	config  *config.Config
	logger  logging.Logger
	auther  *AutherService
}

func (w *WebService) Home(ctx context.Context) (err error) {
	return nil
}

func NewWebService(
	clients *data.Clients,
	cfg *config.Config,
	logger logging.Logger,
	auther *AutherService,
) websvc.Service {

	return &WebService{
		clients: clients,
		logger:  logger,
		auther:  auther,
		config:  cfg,
	}
}

func RegisterWebHTTPService(
	mux goahttp.Muxer,
	clients *data.Clients,
	svcs *services.Services,
	config *config.Config,
	logger logging.Logger,
	auther *AutherService,
) {
	eh := errorHandler(logger)
	svc := NewWebService(clients, config, logger, auther)
	endpoints := websvc.NewEndpoints(svc)
	handler := webhttp.New(
		endpoints, mux, decoder, encoder, eh, nil, nil,
	)

	// handler2 := otelhttp.NewHandler(handler, "auth-service")
	webhttp.Mount(mux, handler)
}

// FileServer provides a static file server with proper handling for SPA routing
func FileServer(rootPath string, router chi.Router) http.Handler {
	fs := http.FileServer(http.Dir("./web/client/dist"))
	router.Handle("/_astro/", fs)
	router.Handle("/assets/", fs)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1/") {
			http.NotFound(w, r)
			return
		}

		// Check if the path exists as a file in the dist directory
		requestedPath := filepath.Join(rootPath, r.URL.Path)
		if r.URL.Path != "/" && !strings.HasPrefix(r.URL.Path, "/_astro/") {
			// Add .html extension for non-root paths that don't have it
			if filepath.Ext(requestedPath) == "" {
				requestedPath += "/index.html"
			}
		}

		// If the file exists, serve it
		if _, err := os.Stat(requestedPath); err == nil {
			http.ServeFile(w, r, requestedPath)
			return
		}

		// For paths that don't exist as files, serve index.html
		// (This allows for client-side routing if you're using it)
		indexPath := filepath.Join(rootPath, "index.html")
		http.ServeFile(w, r, indexPath)
	})
}

// ProtectedFrontendFileServer serves static files with frontend route protection
func ProtectedFrontendFileServer(
	rootPath string,
	router chi.Router,
	protection *middleware.FrontendRouteProtection,
	logger logging.Logger,
) http.Handler {
	fs := http.FileServer(http.Dir(rootPath))

	// Create file server handler
	fileHandler := astro_fs.NewCustomFileHandler(rootPath, logger, &astro_fs.Configuration{
		Directory:     rootPath,
		CacheTTL:      0,
		EnableGzip:    true,
		LogFile:       "",
		Debug:         true,
		SPA:           false,
		IndexFile:     "index.html",
		MaxLogSize:    0,
		MaxLogBackups: 0,
		MaxLogAge:     0,
		AstroMode:     true,
		PreloadAssets: true,
		HTTP2:         true,
	})

	// Configure static file paths to bypass protection
	for _, prefix := range protection.GetStaticFilePrefixes() {
		router.Handle(prefix+"*", fileHandler)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Block direct access to API endpoints through the file server
		if strings.HasPrefix(path, "/v1/") {
			http.NotFound(w, r)
			return
		}

		// For static assets, serve directly without protection
		for _, prefix := range protection.GetStaticFilePrefixes() {
			if strings.HasPrefix(path, prefix) {
				fs.ServeHTTP(w, r)
				return
			}
		}

		// Check authentication for protected routes
		if protection.IsProtectedPath(path) {
			// Check if user is authenticated
			sess, err := utils.GetSession(r, protection.Config())
			authenticated := false

			if err == nil {
				userID, ok := sess.Values["user_id"].(string)
				if ok && userID != "" {
					authenticated = true
				}
			}

			if !authenticated {
				// Redirect to login page with return URL
				returnURL := r.URL.String()
				loginURL := protection.LoginPath()

				if returnURL != loginURL && returnURL != "/" {
					loginURL += "?returnUrl=" + returnURL
				}

				http.Redirect(w, r, loginURL, http.StatusFound)
				return
			}
		}

		// // Check if the path exists as a file in the dist directory
		// requestedPath := filepath.Join(rootPath, path)
		// if path != "/" {
		// 	// Add index.html for directories or paths without extensions
		// 	if stat, err := os.Stat(requestedPath); err == nil && stat.IsDir() {
		// 		requestedPath = filepath.Join(requestedPath, "index.html")
		// 	} else if filepath.Ext(requestedPath) == "" {
		// 		// Try with .html extension
		// 		htmlPath := requestedPath + ".html"
		// 		if _, err := os.Stat(htmlPath); err == nil {
		// 			requestedPath = htmlPath
		// 		} else {
		// 			// Try as directory with index.html
		// 			indexPath := filepath.Join(requestedPath, "index.html")
		// 			if _, err := os.Stat(indexPath); err == nil {
		// 				requestedPath = indexPath
		// 			}
		// 		}
		// 	}
		// } else {
		// 	// For root path, serve index.html
		// 	requestedPath = filepath.Join(rootPath, "index.html")
		// }
		//
		// // If the file exists, serve it
		// if _, err := os.Stat(requestedPath); err == nil {
		// 	// Set the appropriate content type
		// 	contentType := getContentType(requestedPath)
		// 	w.Header().Set("Content-Type", contentType)
		//
		// 	http.ServeFile(w, r, requestedPath)
		// 	return
		// }
		//
		// // For paths that don't exist as files, serve index.html (for SPA routing)
		// logger.Debug("File not found, serving index.html for SPA routing",
		// 	logging.String("path", path),
		// 	logging.String("requested_path", requestedPath))
		//
		// indexPath := filepath.Join(rootPath, "index.html")
		// w.Header().Set("Content-Type", "text/html")
		// w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; connect-src *")
		// http.ServeFile(w, r, indexPath)

		fileHandler.ServeHTTP(w, r)
	})
}

// RegisterFrontendRoutes sets up protected frontend routes
func RegisterFrontendRoutes(
	router chi.Router,
	rootPath string,
	svcs *services.Services,
	config *config.Config,
	logger logging.Logger,
) {
	// Create frontend route protection middleware
	protection := middleware.NewFrontendRouteProtection(
		config,
		logger,
		svcs.Session,
		svcs.CookieHandler,
	)

	// Configure public and protected paths
	protection.SetPublicPaths(
		"/",
		"/login",
		"/register",
		"/signup",
		"/forgot-password",
		"/reset-password",
		"/verify-email",
		"/about",
		"/contact",
		"/legal/*",
		"/_astro/*",
		"/assets/*",
	)

	protection.SetProtectedPaths(
		"/dashboard",
		"/dashboard/*",
		"/profile",
		"/settings",
		"/projects",
		"/projects/*",
		"/reports",
		"/reports/*",
		"/admin",
		"/admin/*",
	)

	authMw := middleware.Auth(config, logger, svcs.Session, svcs.APIKey, false)

	// Add this to your routes
	router.HandleFunc("/debug/log", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var logData struct {
			Level   string `json:"level"`
			Message string `json:"message"`
			Source  string `json:"source"`
		}

		if err := json.NewDecoder(r.Body).Decode(&logData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Log to server console
		logger.Info("Client log",
			logging.String("level", logData.Level),
			logging.String("message", logData.Message),
			logging.String("source", logData.Source))

		w.WriteHeader(http.StatusOK)
	})

	router.Group(func(r chi.Router) {
		r.Use(authMw)
		protection.SetLoginPath("/login")
		protection.SetStaticFilePrefixes("/_astro/", "/assets/")
		// Apply the middleware to all routes
		r.Use(protection.ProtectFrontendRoutes)

		// Set up the file server
		fs := ProtectedFrontendFileServer(rootPath, router, protection, logger)

		r.Handle("/*", fs)
	})
}
