package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/thinhcompany/go-ecommerce-backend-api/internal/controller"
	middleware "github.com/thinhcompany/go-ecommerce-backend-api/internal/middlewares"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	v2 := r.Group("v2/2025")
	{
		v2.GET("/ping", c.NewPongController().PongWithDefaultUID)
		v2.GET("/user/1", c.NewUserController().GetUserByID)
	}

	return r
}
