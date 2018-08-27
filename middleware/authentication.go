package middleware

import (
	"github.com/gin-gonic/gin"
	"qrapi/g/x/web"
	"qrapi/o/auth"
)

const errNotPermision = "you are not enough permision"
const errUnauthorize = "you are not enough permision"

func Authenticate(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = web.GetToken(c.Request)
		var res, err = auth.GetByID(token)
		if err == nil && res != nil {
			if res.Role != role {
				c.AbortWithStatusJSON(401, map[string]interface{}{
					"error":  errNotPermision,
					"status": "error",
				})
				return
			}
		} else {
			if err.Error() == "not found" {
				c.AbortWithStatusJSON(401, map[string]interface{}{
					"error":  errUnauthorize,
					"status": "error",
				})
				return
			}
		}
		c.Next()
	}
}

var MustBeAdmin = Authenticate("admin")
var MustBeBoss = Authenticate("boss")
var MustBeManager = Authenticate("manager")
