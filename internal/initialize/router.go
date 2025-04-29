package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/go-ecommerce-backend-api/global"
	"github.com/thinhcompany/go-ecommerce-backend-api/internal/routers"
	//c "github.com/thinhcompany/go-ecommerce-backend-api/internal/controller"
	//middleware "github.com/thinhcompany/go-ecommerce-backend-api/internal/middlewares"
)

func InitializeRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User
	mainGroup := r.Group("v1/2024")
	{
		mainGroup.GET("/checkStatus")

		//
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)

		//
		managerRouter.InitAdminRouter(mainGroup)
		managerRouter.InitUserRouter(mainGroup)
	}

	return r
}

// func InitializeRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.Use(middleware.AuthMiddleware())

// 	v2 := r.Group("v2/2025")
// 	{
// 		v2.GET("/ping", c.NewPongController().PongWithDefaultUID)
// 		v2.GET("/user/1", c.NewUserController().GetUserByID)
// 	}

// 	return r
// }
