package order

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
)

const (
	PENDING_STATE  = "pending"
	DELIVERY_STATE = "delivery"
	DONE_STATE     = "done"
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
	Activated     bool   `bson:"activated" json:"activated"`
}

func (order *Order) Create() error {
	order.Activated = false
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

func DeliveryOrder(id string) error {
	return orderTable.UpdateId(id, bson.M{
		"$set": bson.M{
			"activated": true,
		},
	})
}

func GetOrderByID(id string) (*Order, error) {
	var order *Order
	var err = orderTable.FindID(id, &order)
	return order, err
}
