package customer

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi/g/x/web"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
	"qrapi/x/validator"
)

var customerTable = mongodb.NewTable("customer", "cus")
var customerLog = logger.NewLogger("customer")

const (
	CUSTOMER = "customer"
	ADMIN    = "admin"
)

var (
	errExistEmail = "Xin lỗi! Email đã tồn tại trong hệ thống"
	errExistPhone = "Xin lỗi! Số điện thoại đã tồn tại trong hệ thống"
)

type Customer struct {
	mongodb.Model  `bson:",inline"`
	Name           string   `bson:"name" json:"name"`
	Phone          string   `bson:"phone" json:"phone"`
	Email          string   `bson:"email" json:"email"`
	HashedPassword string   `bson:"password" json:"-"`
	Password       Password `bson:"-" json:"password"`
	CompanyName    string   `bson:"company_name" json:"company_name"`
	Logo           string   `bson:"logo" json:"logo"`
	Information    string   `bson:"information" json:"information"`
	Role           string   `bson:"role" json:"role"`
}

const (
	errExists           = "user exists!"
	errMisMatchUNamePwd = "username or password is incorect!"
)

func (u *Customer) CreateAccount() error {
	if user, _ := GetCustomerByPhone(u.Phone); user != nil {
		return web.BadRequest(errExists)
	}
	return customerTable.Create(u)
}

func GetCustomerByPhone(phone string) (*Customer, error) {
	var customer *Customer
	var err = customerTable.FindOne(bson.M{
		"phone": phone,
	}, &customer)
	return customer, err
}

func (u *Customer) Create() error {
	var err = validator.Struct(u)
	hashed, _ := u.Password.Hash()
	u.HashedPassword = hashed
	if err != nil {
		customerLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	var existEmail, _ = GetCustomerByEmail(u.Email)
	if existEmail != nil {
		return web.BadRequest(errExistEmail)
	}
	var existPhone, _ = GetCustomerByPhone(u.Phone)
	if existPhone != nil {
		return web.BadRequest(errExistPhone)
	}
	return customerTable.Create(u)
}

func GetAdmin(uname string, role string) (*Customer, error) {
	var customer *Customer
	var err = customerTable.FindOne(bson.M{
		"uname": uname,
		"role":  role,
	}, &customer)
	return customer, err
}

func GetCustomerByEmail(email string) (*Customer, error) {
	var customer *Customer
	var err = customerTable.FindOne(bson.M{
		"email": email,
		"role":  CUSTOMER,
	}, &customer)
	return customer, err
}

func DeleteUserByID(id string) error {
	return customerTable.DeleteByID(id)
}
