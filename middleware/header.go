package middleware

import (
	"github.com/gin-gonic/gin"
)

func AddHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Range, Content-Disposition, Authorization",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Credentials",
			"true",
		)
		//remember
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(200)
			return
		}
		c.Next()
	}
}
