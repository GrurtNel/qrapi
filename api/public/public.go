package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"qrapi/g/x/web"
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
	s.GET("product/detail", s.getProduct)
	return &s
}

func (s *PublicServer) scanQrcode(c *gin.Context) {
	var token = c.Query("token")
	var err, str = fcm.SendToOne(token, fcm.FmcMessage{Title: "Hello", Body: "Anh"})
	fmt.Println(err)
	fmt.Println(str)
	s.Success(c)
}

const CIPHER_KEY = "trungyeulen4ever"

//CTwHkYUk2AdRqEkLbfa6_rlK6_-x6h78ySbTNs1k4eDu9Zj1DtWlRg==
func (s *PublicServer) getProduct(c *gin.Context) {
	var id = c.Query("id")
	decrypted, err := security.Decrypt([]byte(CIPHER_KEY), id)
	web.AssertNil(err)
	glog.Info(decrypted)
	s.SendData(c, map[string]interface{}{
		"product": "",
		"ds":      "",
	})
}

func getCustomerProductID(gid string) (string, string) {
	return strings.Split(gid, "$$")[0], strings.Split(gid, "$$")[1]
}
