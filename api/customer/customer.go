package customer

import (
	"github.com/gin-gonic/gin"
	"qrapi/g/x/web"
	"qrapi/middleware"
	"qrapi/o/product"
)

type CustomerServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewCustomerServer(parent *gin.RouterGroup, name string) *CustomerServer {
	var s = CustomerServer{
		RouterGroup: parent.Group(name),
	}
	s.Use(middleware.MustBeCustomer)
	s.POST("product/create", s.createProduct)
	s.GET("product/list", s.getProducts)
	return &s
}

func (s *CustomerServer) createProduct(c *gin.Context) {
	var product *product.Product
	c.BindJSON(&product)
	web.AssertNil(product.Create())
	s.SendData(c, product)
}

func (s *CustomerServer) getProducts(c *gin.Context) {
	var customerID = c.Query("user_id")
	var products, err = product.GetProductsByCustomer(customerID)
	web.AssertNil(err)
	s.SendData(c, products)
}
