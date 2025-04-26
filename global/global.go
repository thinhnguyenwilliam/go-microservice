package global

import (
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/logger"
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Mdb    *gorm.DB
	Logger *logger.LoggerZap
)
