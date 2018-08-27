package fcm

import (
	fcm "github.com/NaySoftware/go-fcm"
)

const (
	RESPONSE_FAIL = "fail"
)

type FcmClient struct {
	*fcm.FcmClient
}

type FmcMessage struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

var f = &FcmClient{}

func NewFCM(serverKey string) *FcmClient {
	f.FcmClient = fcm.NewFcmClient(serverKey)
	return &FcmClient{
		FcmClient: fcm.NewFcmClient(serverKey),
	}
}

func SendToMany(ids []string, data FmcMessage) (error, string) {
	var noti = fcm.NotificationPayload{
		Title: data.Title,
		Body:  data.Body,
		Sound: "ting.wav",
	}
	f.NewFcmRegIdsMsg(ids, data)
	f.SetNotificationPayload(&noti)
	status, err := f.Send()
	if err != nil {
		return err, RESPONSE_FAIL
	}
	return nil, status.Err
}

func SendToOne(id string, data FmcMessage) (error, string) {
	return SendToMany([]string{id}, data)
}
