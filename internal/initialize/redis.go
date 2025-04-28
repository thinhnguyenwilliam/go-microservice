package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/thinhcompany/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

var (
	Ctx = context.Background()
)

func InitializeRedis() {
	// Read Redis configuration from global.Config
	redisConfig := global.Config.Redis
	redisAddr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)

	// Initialize Redis client
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,            // Redis address from config
		Password: redisConfig.Password, // Redis password from config
		DB:       redisConfig.DB,       // Redis DB index from config
		PoolSize: 10,                   // Connection pool size
	})

	// Test the connection
	pong, err := global.Rdb.Ping(Ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	fmt.Println("Connected to Redis:", pong)

	// Call redisExample after successful Redis initialization
	redisExample()
}

// Example function to interact with Redis
func redisExample() {
	// Set a value in Redis with a key "score"
	err := global.Rdb.Set(Ctx, "score", "100", 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set key in Redis", zap.Error(err))
	} else {
		fmt.Println("Key 'score' set to 100 in Redis.")
		global.Logger.Info("Successfully set 'score' key in Redis", zap.String("score", "100"))
	}

	// Get the value of "score" from Redis
	value, err := global.Rdb.Get(Ctx, "score").Result()
	if err != nil {
		if err == redis.Nil {
			// Redis key doesn't exist
			global.Logger.Warn("Key 'score' does not exist in Redis")
		} else {
			global.Logger.Error("Failed to get key from Redis", zap.Error(err))
		}
	} else {
		// Log the retrieved value
		fmt.Println("Value of 'score' in Redis:", value)
		global.Logger.Info("Retrieved value from Redis", zap.String("score", value))
	}
}
