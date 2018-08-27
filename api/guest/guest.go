package guest

import (
	"github.com/gin-gonic/gin"
	"qrapi/g/x/web"
	"qrapi/o/auth"
	"qrapi/o/customer"
)

type GuestServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewGuestServer(parent *gin.RouterGroup, name string) *GuestServer {
	var s = GuestServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("register", s.register)
	return &s
}

func (s *GuestServer) register(c *gin.Context) {
	var u *customer.Customer
	c.BindJSON(&u)
	u.Role = customer.CUSTOMER
	web.AssertNil(u.Create())
	var auth = auth.Create(u.ID, u.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": u,
	})
	s.SendData(c, u)
}
