package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhcompany/go-ecommerce-backend-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer your-secret-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": response.ErrorInvalidToken,
				"msg":  response.Msg(response.ErrorInvalidToken),
			})
			return
		}
		c.Next()
	}
}
