package middleware

import (
	"github.com/gin-gonic/gin"
)

func AddStaticHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/x-javascript")
	}
}
