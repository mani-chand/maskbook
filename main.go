package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	// We no longer need "strings" for the old FileServer
	// "strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()

	// === Middleware ===
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set up CORS
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

	staticDir := "frontend/dist"

	// Create a file server handler for the static directory
	fileServer := http.FileServer(http.Dir(staticDir))

	// Use a custom handler to serve files and fallback to index.html
	// This is the new, correct logic for a Single Page App (SPA)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Get the path to the requested file
		filePath := filepath.Join(staticDir, r.URL.Path)

		// Check if a file exists at that path
		// e.g., /assets/index-123.js
		stat, err := os.Stat(filePath)

		// If the file exists and is not a directory, serve it
		if err == nil && !stat.IsDir() {
			fileServer.ServeHTTP(w, r)
			return
		}

		// If the file does NOT exist (e.g., /create, /login),
		// serve the index.html
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

// The old, problematic FileServer function has been removed.
