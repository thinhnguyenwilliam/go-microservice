package main

import (
	"github.com/gin-contrib/cors"
	"github.com/thinhcompany/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.SetupRouter()

	// Use default CORS settings
	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://yourdomain.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}))

	r.Run(":8080") // Start server on port 8080
}
