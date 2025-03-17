package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Grouping routes under /v1/2025
	v1 := r.Group("v1/2025")
	{
		v1.GET("/ping", Pong) // curl http://localhost:8080/v1/2025/ping
	}

	v2 := r.Group("v2/2025")
	{
		v2.GET("/ping/:name", PongWithName) // curl http://localhost:8080/v2/2025/ping/John
		//v2.GET("/ping", PongWithUID)        // curl http://localhost:8080/v2/2025/ping?uid=12345
		v2.GET("/ping", PongWithDefaultUID) // curl http://localhost:8080/v2/2025/ping?uid=12345
	}

	return r
}

// Handler for /v2/2025/ping?uid=12345 (query parameter with default value)
func PongWithDefaultUID(c *gin.Context) {
	uid := c.DefaultQuery("uid", "guest") // If 'uid' is missing, default to "guest"

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, your UID is " + uid + "!",
	})
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Gin!",
	})
}

// Handler for /v2/2025/ping/:name (extracts name parameter)
func PongWithName(c *gin.Context) {
	name := c.Param("name") // Extract 'name' from the URL
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + name + "!",
	})
}

// Handler for /v2/2025/ping?uid=12345 (query parameter)
func PongWithUID(c *gin.Context) {
	uid := c.Query("uid") // Get 'uid' query parameter
	if uid == "" {
		uid = "unknown" // Default value if 'uid' is not provided
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, your UID is " + uid + "!",
	})
}
