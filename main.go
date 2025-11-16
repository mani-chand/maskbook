package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	// 1. Import your local 'services' package
	// Make sure "noob" is your module name from go.mod
	"noob/services/post"
	services "noob/services/user"
)

// 2. The 'authInput' struct has been moved to services/user.go

func main() {
	// gin.Default() comes with Logger and Recoverer middleware
	r := gin.Default()

	// Set up CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// === API Routes ===
	// Mount your API routes under /api
	api := r.Group("/api")
	{
		// I've added GetPostsHandler back in so this route works
		// 3. Call the EXPORTED (capitalized) functions
		api.POST("/createUser", services.CreateUser)
		api.POST("/login", services.ValidateUser)
		api.GET("/posts", post.Get_all_posts)
		api.POST("/post", post.Create_post)
		// r.Get("/user/{userID}", GetUserHandler)
	}

	// === Frontend (Svelte) Serving ===
	staticDir := "frontend/dist"

	// Check if static dir exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory %s does not exist. Run 'npm run build' in the frontend directory.", staticDir)
	}

	// This middleware serves files from the static dir.
	// If a file is found (e.g., /assets/index-123.js), it serves it.
	// If not, it passes control to the next handler (e.g., NoRoute).
	r.Use(static.Serve("/", static.LocalFile(staticDir, false)))

	// Fallback for SPA (Single Page App)
	// If no file or API route is matched, serve index.html
	r.NoRoute(func(c *gin.Context) {
		// Don't serve index.html for API calls that 404
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{"error": "api_route_not_found"})
			return
		}

		// Serve the index.html
		indexPath := filepath.Join(staticDir, "index.html")
		if _, err := os.Stat(indexPath); os.IsNotExist(err) {
			log.Printf("index.html not found in %s", staticDir)
			c.JSON(http.StatusNotFound, gin.H{"error": "index_not_found"})
			return
		}
		c.File(indexPath)
	})

	log.Println("Starting server on :8080...")
	// Use r.Run() which is Gin's equivalent of http.ListenAndServe
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

// GetPostsHandler is a sample API handler, updated for Gin
// (I've added this back from our previous version so the route works)
func GetPostsHandler(c *gin.Context) {
	// In a real app, you'd fetch this from your database
	response := `[
		{"id": "1", "message": "Hello from your Go backend!"},
		{"id": "2", "message": "Svelte and Go are a great combo!"}
	]`

	// c.Data sends raw bytes. It's equivalent to w.Write()
	c.Data(http.StatusOK, "application/json", []byte(response))
}
