package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error_code":    0,
		"error_message": "success",
	})
}
