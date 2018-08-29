package product

import (
	"qrapi/g/x/web"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
	"qrapi/x/validator"
)

var productLog = logger.NewLogger("tbl_product")
var productTable = mongodb.NewTable("post", "prd")

type Product struct {
	mongodb.Model `bson:",inline"`
	Title         string `bson:"title" json:"title" validate:"string"`
	Content       string `bson:"content" json:"content" validate:"string,min=0"`
	Description   string `bson:"description" json:"description" validate:"string,min=0"`
	Category      string `bson:"category" json:"category" validate:"string,min=0"`
	Author        string `bson:"author" json:"author"`
	Approve       bool   `bson:"approve" json:"approve"`
}

func (product *Product) Create() error {
	err := validator.Struct(product)
	if err != nil {
		productLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return productTable.Create(product)
}
