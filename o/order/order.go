package order

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
)

const (
	PENDING_STATE     = "pending"
	PROGRESSING_STATE = "progressing"
	DONE_STATE        = "done"
)

var orderLog = logger.NewLogger("tbl_order")
var orderTable = mongodb.NewTable("order", "prd")

type Order struct {
	mongodb.Model `bson:",inline"`
	Name          string `bson:"name" json:"name"`
	Type          string `bson:"type" json:"type"`
	CustomerID    string `bson:"customer_id" json:"customer_id"`
	ProductID     string `bson:"product_id" json:"product_id"`
	Quantity      int    `bson:"quantity" json:"quantity"`
	URL           string `bson:"url" json:"url"`
	Status        string `bson:"status" json:"status"`
}

func (order *Order) Create() error {
	order.Status = PENDING_STATE
	return orderTable.Create(order)
}

func GetOrdersByCustomer(customerID string) ([]*Order, error) {
	glog.Info(customerID)
	var orders []*Order
	var err = orderTable.FindWhere(bson.M{"customer_id": customerID}, &orders)
	return orders, err
}

func GetOrders() ([]*Order, error) {
	var order []*Order
	var err = orderTable.FindAll(&order)
	return order, err
}

func GetOrderByID(id string) (*Order, error) {
	var order *Order
	var err = orderTable.FindID(id, &order)
	return order, err
}
