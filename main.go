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
	"github.com/joho/godotenv" // 1. Import godotenv

	// 2. THIS IS THE CORRECT IMPORT.
	"noob/services/database"

	"noob/services/post"
	"noob/services/storage"
	"noob/services/user"
)

func loadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func main() {
	// 3. Load environment variables first
	loadEnv()

	// 4. Connect to the database using the correct package
	database.ConnectDB()

	// 5. Initialize S3 Client
	storage.InitS3()

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
		// Call the functions from their correct packages
		api.GET("/posts", post.Get_all_posts)
		api.POST("/createUser", user.CreateUser)
		api.POST("/login", user.ValidateUser)
		api.POST("/createPost", post.Create_post)
		// You will add your post creation route here later
		// e.g., api.POST("/createPost", post.CreatePost)
	}

	// === Frontend (SVvelte) Serving ===
	staticDir := "frontend/dist"

	// Check if static dir exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory %s does not exist. Run 'npm run build' in the frontend directory.", staticDir)
	}

	// This middleware serves files from the static dir.
	r.Use(static.Serve("/", static.LocalFile(staticDir, false)))

	// Fallback for SPA (Single Page App)
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
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
