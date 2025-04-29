package manager

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// Dummy middleware placeholders
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement rate limiting
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement authentication check
		c.Next()
	}
}

func PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement permission check
		c.Next()
	}
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private admin routes with middleware
	userRouterPrivate := Router.Group("/admin/user",
		RateLimiter(),
		AuthMiddleware(),
		PermissionMiddleware(),
	)
	{
		userRouterPrivate.POST("/active_user")
	}
}
