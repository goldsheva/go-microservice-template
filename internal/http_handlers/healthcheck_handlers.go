package http_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHttpHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
