package auth

import (
	"github.com/gin-gonic/gin"
	"qrapi/g/x/web"
	"qrapi/o/auth"
	"qrapi/o/customer"
)

type AuthServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAuthServer(parent *gin.RouterGroup, name string) *AuthServer {
	var s = AuthServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("login", s.login)
	s.GET("super-admin", s.checkSuperAdmin)
	return &s
}

func (s *AuthServer) login(c *gin.Context) {
	var loginInfo = struct {
		Uname    string
		Password string
	}{}
	c.BindJSON(&loginInfo)
	user, err := customer.Login(loginInfo.Uname, loginInfo.Password)
	web.AssertNil(err)
	var auth = auth.Create(user.ID, user.Role)
	s.SendData(c, map[string]interface{}{
		"token":     auth.ID,
		"user_info": user,
	})
}

func (s *AuthServer) checkSuperAdmin(c *gin.Context) {
	var superAdmin, err = customer.GetAdmin("", "super-admin")
	web.AssertNil(err)
	s.SendData(c, superAdmin)
}
