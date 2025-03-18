package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/go-ecommerce-backend-api/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc==> user controller
// us==> user service
// controller --> service --> repo --> models --> db
func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"cr7", "m10", "william"},
	})
}
