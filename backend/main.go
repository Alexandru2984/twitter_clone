package main

import (
	"log"
	"os"

	"github.com/Alexandru2984/twitter_clone/backend/config"
	"github.com/Alexandru2984/twitter_clone/backend/middleware"
	"github.com/Alexandru2984/twitter_clone/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Setup router
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORSMiddleware())

	// Public routes
	api := router.Group("/api")
	{
		// Auth routes
		api.POST("/register", routes.Register)
		api.POST("/login", routes.Login)

		// Public tweet routes
		api.GET("/tweets", routes.GetTweets)
		api.GET("/tweets/user/:username", routes.GetUserTweets)

		// Public user routes
		api.GET("/users/:username", routes.GetUserProfile)
	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Auth routes
		protected.GET("/profile", routes.GetProfile)

		// Tweet routes
		protected.POST("/tweets", routes.CreateTweet)
		protected.POST("/tweets/:id/like", routes.LikeTweet)
		protected.DELETE("/tweets/:id", routes.DeleteTweet)

		// User routes
		protected.POST("/users/:username/follow", routes.FollowUser)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}
