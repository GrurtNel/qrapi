package auth

import (
	"qrapi/x/mongodb"
)

var authTable = mongodb.NewTable("auth", "auth")

type Auth struct {
	mongodb.Model `bson:",inline"`
	UserID        string `bson:"user_id" json:"user_id"`
	Role          string `bson:"role" json:"role"`
	Revoked       bool   `bson:"revoked" json:"revoked"`
}

func Create(userID, role string) *Auth {
	var auth = &Auth{
		UserID:  userID,
		Role:    role,
		Revoked: false,
	}
	authTable.CreateAuth(auth)
	return auth
}

func GetByID(id string) (*Auth, error) {
	var auth *Auth
	return auth, authTable.FindID(id, &auth)
}

func DeleteByID(id string) error {
	return authTable.RemoveId(id)
}
