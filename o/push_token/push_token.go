package push_token

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi/x/mongodb"
)

var pushTable = mongodb.NewTable("push_token", "push")

type PushToken struct {
	PushToken string `bson:"push_token" json:"push_token" `
	DeviceID  string `bson:"device_id" json:"device_id" `
}

func (push *PushToken) Create() error {
	var _, err = pushTable.Upsert(bson.M{
		"device_id": push.DeviceID,
	}, push)
	return err
}

func GetAllPushToken() ([]string, error) {
	var push []*PushToken
	var result []string
	var err = pushTable.FindAll(&push)
	if push != nil {
		for _, item := range push {
			result = append(result, item.PushToken)
		}
	}
	return result, err
}

// func GetByDeviceID(deviceID string) (*PushToken, error) {
// 	var push *PushToken
// 	return push, pushTable.Find(bson.M{
// 		"device_id": deviceID,
// 	}).One(&push)
// }
