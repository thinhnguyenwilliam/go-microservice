package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// unit test: https://chatgpt.com/c/67d7fb39-671c-8005-87b1-03b0af2924ff
type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

// Handler for /v2/2025/ping?uid=12345 (query parameter with default value)
func (p *PongController) PongWithDefaultUID(c *gin.Context) {
	uid := c.DefaultQuery("uid", "guest") // If 'uid' is missing, default to "guest"

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, your UID is " + uid + "!",
	})
}
