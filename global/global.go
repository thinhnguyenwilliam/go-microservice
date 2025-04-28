package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/logger"
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Mdb    *gorm.DB
	Logger *logger.LoggerZap
	Rdb    *redis.Client
)
