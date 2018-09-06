package public

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"qrapi/common"
	"qrapi/g/x/web"
	"qrapi/o/admin"
	"qrapi/o/customer"
	"qrapi/o/order"
	"qrapi/o/product"
	sHistory "qrapi/o/scan_history"
	"qrapi/x/fcm"
	"qrapi/x/security"
	"strings"
)

type PublicServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewPublicServer(parent *gin.RouterGroup, name string) *PublicServer {
	var s = PublicServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("qrcode/scan", s.scanQrcode)
	s.GET("marketing/scan", s.scanMarketing)
	s.GET("product/detail", s.getProduct)
	s.GET("order/detail", s.getOrder)
	s.POST("register", s.register)
	return &s
}

func (s *PublicServer) scanMarketing(c *gin.Context) {
	var scanHistory = &sHistory.ScanHistory{
		OrderID: c.Query("order_id"),
	}
	scanHistory.SetID(c.Query("order_id"))
	web.AssertNil(scanHistory.Create())
	s.Success(c)
}

func (s *PublicServer) scanQrcode(c *gin.Context) {
	var token = c.Query("token")
	var err, str = fcm.SendToOne(token, fcm.FmcMessage{Title: "Hello", Body: "Anh"})
	fmt.Println(err)
	fmt.Println(str)
	s.Success(c)
}

func (s *PublicServer) register(c *gin.Context) {
	var admin *admin.Admin
	web.AssertNil(c.BindJSON(&admin))
	web.AssertNil(admin.CreateAccount())
	s.Success(c)
}

func (s *PublicServer) getOrder(c *gin.Context) {
	var order, err = order.GetOrderByID(c.Query("order_id"))
	web.AssertNil(err)
	s.SendData(c, order)
}

//CTwHkYUk2AdRqEkLbfa6_rlK6_-x6h78ySbTNs1k4eDu9Zj1DtWlRg==
var errNotValidCode = errors.New("Không tìm thấy thông tin sản phẩm")

func (s *PublicServer) getProduct(c *gin.Context) {
	var id = c.Query("id")
	var code = c.Query("code")
	if code != "" {
		id = id + code
	}
	decrypted, err := security.Decrypt([]byte(common.CIPHER_KEY), id)
	if err != nil {
		web.AssertNil(errNotValidCode)
	}
	customerID, productID := getCustomerProductID(decrypted)
	customer, err := customer.GetCustomerByID(customerID)
	web.AssertNil(err)
	if customer == nil {
		web.AssertNil(errNotValidCode)
	}
	product, err := product.GetProductByID(productID)
	web.AssertNil(err)
	if product == nil {
		web.AssertNil(errNotValidCode)
	}
	s.SendData(c, map[string]interface{}{
		"product":  product,
		"customer": customer,
	})
}

func getCustomerProductID(gid string) (string, string) {
	return strings.Split(gid, "$$")[0], strings.Split(gid, "$$")[1]
}
