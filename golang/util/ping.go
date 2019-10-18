package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(e *gin.Engine, path string) {
	e.GET(path, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"error_code":    0,
			"error_message": "success",
		})
	})
}
