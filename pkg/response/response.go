package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int    `json:"code"`    // HTTP status code
	Message string `json:"message"` // Message to be displayed to the user
	Data    any    `json:"data"`    // Data to be returned to the user
}

// success response
func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// error response
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusBadGateway, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
