package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/go-ecommerce-backend-api/internal/service"
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/response"
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
	// response.SuccessResponse(c, response.ErrorCodeSuccess, gin.H{
	// 	"info":  uc.userService.GetInfoUser(),
	// 	"users": []string{"cr7", "m10", "william1"},
	// })

	response.ErrorResponse(c, response.ErrorCodeParamInvalid)

}
