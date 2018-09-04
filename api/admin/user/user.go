package user

import (
	"github.com/gin-gonic/gin"
	"qrapi/g/x/web"
	"qrapi/o/customer"
)

type UserServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewUserServer(parent *gin.RouterGroup, name string) *UserServer {
	var s = UserServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("list", s.getUsers)
	s.POST("create", s.createUser)
	s.GET("delete", s.deleteUser)
	return &s
}

func (s *UserServer) getUsers(c *gin.Context) {
	var users, err = customer.GetUsers()
	web.AssertNil(err)
	s.SendData(c, users)
}

func (s *UserServer) createUser(c *gin.Context) {
	var customer *customer.Customer
	c.BindJSON(&customer)
	web.AssertNil(customer.Create())
	s.SendData(c, customer)
}

func (s *UserServer) deleteUser(c *gin.Context) {
	var id = c.Query("id")
	web.AssertNil(customer.DeleteUserByID(id))
	s.Success(c)
}
