package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()

	// === Middleware ===
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set up CORS.
	// This is good practice for your API, even in production.
	// You can be more restrictive here if you want.
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})
	r.Use(corsMiddleware.Handler)

	// === API Routes ===
	// Mount your API routes under /api
	r.Route("/api", func(r chi.Router) {
		r.Get("/posts", GetPostsHandler)
		// r.Post("/posts", CreatePostHandler)
		// r.Get("/user/{userID}", GetUserHandler)
	})

	// === Frontend (Svelte) Serving ===
	// Set the path to your built Svelte app
	// Updated to 'frontend/dist' to match your plain Vite setup
	staticDir := "frontend/dist" 
	
	// 1. Serve static files (JS, CSS, images, etc.)
	FileServer(r, "/", http.Dir(staticDir))

	// 2. Serve the index.html for any other route to support
	// Svelte's client-side routing
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

// GetPostsHandler is a sample API handler
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	// In a real app, you'd fetch this from your database
	response := `[
		{"id": "1", "message": "Hello from your Go backend!"},
		{"id": "2", "message": "Svelte and Go are a great combo!"}
	]`

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

// FileServer conveniently serves static files and handles SPA (Single Page App)
// routing.
func FileServer(r chi.Router, public string, static http.Dir) {
	if strings.ContainsAny(public, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	root, _ := filepath.Abs(string(static))
	if _, err := os.Stat(root); os.IsNotExist(err) {
		log.Printf("Static directory %s does not exist. Your Svelte app may not be built.", root)
	}

	fs := http.StripPrefix(public, http.FileServer(http.Dir(root)))

	if public != "/" && public[len(public)-1] != '/' {
		// Use the explicit 301 status code instead of the constant
		r.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
		public += "/"
	}
	public += "*"

	r.Get(public, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}