package product

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi/g/x/web"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
	"qrapi/x/validator"
)

var productLog = logger.NewLogger("tbl_product")
var productTable = mongodb.NewTable("product", "prd")

type Product struct {
	mongodb.Model `bson:",inline"`
	Title         string   `bson:"title" json:"title"`
	Gallery       []string `bson:"gallery" json:"gallery"`
	Description   string   `bson:"description" json:"description"`
	CustomerID    string   `bson:"customer_id" json:"customer_id"`
}

func (product *Product) Create() error {
	err := validator.Struct(product)
	if err != nil {
		productLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return productTable.Create(product)
}

func GetProductsByCustomer(customerID string) ([]*Product, error) {
	var products []*Product
	var err = productTable.FindWhere(bson.M{"customer_id": customerID}, &products)
	return products, err
}
