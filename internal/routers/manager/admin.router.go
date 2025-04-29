package manager

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (u *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public routes
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// Private routes with middleware
	adminRouterPrivate := Router.Group("/admin/user")
	// RateLimiter(),
	// AuthMiddleware(),
	// PermissionMiddleware(),

	{
		adminRouterPrivate.POST("/active_user")
	}
}
