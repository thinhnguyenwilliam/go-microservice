package main

import (
	"github.com/thinhcompany/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080") // Start server on port 8080
}
