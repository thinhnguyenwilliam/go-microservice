package initialize

import (
	"fmt"
	"log"

	"github.com/thinhcompany/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Initialize() {
	// Load the configuration
	LoadConfig()
	// Check if DB name is present
	if global.Config.MySQL.DBName == "" {
		log.Fatal("Database name is empty in config")
	}

	// Print loaded configuration info
	fmt.Println("Configuration loaded successfully")
	fmt.Println("Server Port:", global.Config.Server.Port)
	fmt.Println("Database Name:", global.Config.MySQL.DBName)

	// Initialize the logger
	InitializeLogger()
	global.Logger.Info("Logger initialized successfully", zap.String("status", "ok"))

	// Initialize the MySQL database connection
	InitializeMySQL()
	// Initialize the Redis connection

	InitializeRedis()
	global.Logger.Info("Redis initialized successfully", zap.String("status", "ok"))

	r := InitializeRouter()
	r.Run(":8002") // listen and serve on
}
