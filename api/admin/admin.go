package admin

import (
	"github.com/gin-gonic/gin"
	"qrapi/api/admin/category"
	"qrapi/api/admin/customer"
	"qrapi/api/admin/post"
	"qrapi/g/x/web"
	"qrapi/middleware"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeAdmin)
	post.NewPostServer(s.RouterGroup, "post")
	category.NewCategoryServer(s.RouterGroup, "category")
	customer.NewUserServer(s.RouterGroup, "user")
	return &s
}
