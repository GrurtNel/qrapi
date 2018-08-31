package admin

import (
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"qrapi/common"
	"qrapi/g/x/web"
	"qrapi/o/order"
	"qrapi/x/security"
	"strconv"
)

type AdminServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewAdminServer(parent *gin.RouterGroup, name string) *AdminServer {
	var s = AdminServer{
		RouterGroup: parent.Group(name),
	}
	// s.Use(middleware.MustBeAdmin)
	s.GET("order/list", s.getOrders)
	s.GET("order/generate", s.generateSV)
	return &s
}

func (s *AdminServer) getOrders(c *gin.Context) {
	var userID = c.Query("user_id")
	s.SendData(c, userID)
}

func (s *AdminServer) generateSV(c *gin.Context) {
	var numberOfCodes, _ = strconv.Atoi(c.Query("quantity"))
	var orderID = c.Query("id")
	var order, err = order.GetOrderByID(orderID)
	web.AssertNil(err)
	record := []string{"Link sản phẩm", "Mã thẻ cào"}
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	if order.Type == common.QRCOODE_MARKETING {
		for i := 0; i < numberOfCodes; i++ {
			record = []string{order.URL, ""}
			wr.Write(record)
		}
	} else if order.Type == common.QRCOODE_TYPE1 {
		for i := 0; i < numberOfCodes; i++ {
			record = []string{order.URL, ""}
			wr.Write(record)
		}
	} else {
		for i := 0; i < numberOfCodes; i++ {
			var encrypted, _ = security.Encrypt([]byte(common.CIPHER_KEY), order.CustomerID+"$$"+order.ProductID)
			record = []string{encrypted[:len(encrypted)-6], encrypted[len(encrypted)-6 : len(encrypted)]}
			wr.Write(record)
		}
	}

	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	c.Writer.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv

	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
	c.Writer.Write(b.Bytes())
}
