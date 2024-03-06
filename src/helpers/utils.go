package helpers

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/fwg17-cinematix-be/src/services"
)

// KHUSUS ERROR HANDLING
func Utils(err error, ms string, c *gin.Context) {
	if strings.HasPrefix(err.Error(), "sql: no rows") {
		c.JSON(http.StatusNotFound, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	}
}