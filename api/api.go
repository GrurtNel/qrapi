package api

import (
	"qrapi/api/admin"
	"qrapi/api/auth"
	"qrapi/api/guest"
	"qrapi/api/public"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	admin.NewAdminServer(root, "admin")
	auth.NewAuthServer(root, "auth")
	public.NewPublicServer(root, "public")
	guest.NewGuestServer(root, "guest")
}
