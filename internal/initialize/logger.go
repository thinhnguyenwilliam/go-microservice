package initialize

import (
	"github.com/thinhcompany/go-ecommerce-backend-api/global"
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/logger"
)

func InitializeLogger() {
	global.Logger = logger.NewLoggerZap(global.Config.Log)
}
