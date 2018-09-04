package admin

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi/g/x/web"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
)

var adminTable = mongodb.NewTable("admin", "adm")
var adminLog = logger.NewLogger("admin")

const (
	CUSTOMER = "customer"
	ADMIN    = "admin"
)

var (
	errExistEmail                = "Xin lỗi! Email đã tồn tại trong hệ thống"
	errExistPhone                = "Xin lỗi! Số điện thoại đã tồn tại trong hệ thống"
	ErrMismatchedHashAndPassword = "crypto/bcrypt: hashedPassword is not the hash of the given password"
)

type Admin struct {
	mongodb.Model  `bson:",inline"`
	Phone          string   `bson:"phone" json:"phone"`
	HashedPassword string   `bson:"password" json:"-"`
	Password       Password `bson:"-" json:"password"`
	Name           string   `bson:"name" json:"name"`
}

func (u *Admin) CreateAccount() error {
	hashed, _ := u.Password.Hash()
	u.HashedPassword = hashed
	return adminTable.Create(u)
}

func Login(phone, pwd string) (*Admin, error) {
	var admin *Admin
	var query = bson.M{"phone": phone}
	err := adminTable.FindOne(query, &admin)
	if err != nil {
		if err.Error() == "not found" {
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	if err := Password(pwd).ComparePassword(admin.HashedPassword); err != nil {
		if err.Error() == ErrMismatchedHashAndPassword {
			adminLog.Error(err)
			return nil, web.BadRequest("Sai tên đăng nhập hoặc mật khẩu")
		}
		return nil, err
	}
	return admin, nil
}
