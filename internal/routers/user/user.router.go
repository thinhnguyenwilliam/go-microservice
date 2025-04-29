package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routes
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	// private routes
	userRouterPrivate := Router.Group("/user") // or Router.Group("/user", AuthMiddleware())
	{
		userRouterPrivate.GET("/get_info")
	}
}
