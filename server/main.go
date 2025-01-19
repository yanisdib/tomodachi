package main

import (
	"log"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yanisdib/tomodachi/server/config"
	"github.com/yanisdib/tomodachi/server/platform"
)

func main() {
	var once sync.Once

	// Start the server only once
	once.Do(func() {
		router := gin.Default()

		// Setting up CORS policy
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5050"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		}))

		dbPool := config.OpenDBConnectionPool()
		platform.NewPlatformRoutes(router, dbPool)

		if err := router.Run(":5050"); err != nil {
			log.Fatal("An error occured while starting the server at localhost:5050")
		}
	})
}
