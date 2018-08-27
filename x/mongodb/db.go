package mongodb

import (
	"qrapi/x/logger"

	"gopkg.in/mgo.v2"
)

var database *mgo.Database
var mongoLog = logger.NewLogger("mongodb")

func Connect(url string, name string, uname string, pwd string) {
	var info = &mgo.DialInfo{
		Addrs:    []string{url},
		Database: name,
		Username: uname,
		Password: pwd,
	}
	var session *mgo.Session
	var err error
	if uname != "" {
		// session, err = mgo.DialWithInfo(info)
		session, err = mgo.Dial(info.Addrs[0])

	} else {
		session, err = mgo.Dial(url)
	}
	if err != nil {
		mongoLog.Errorf("error connect mongodb %s", err)
		panic(err)
	} else {
		mongoLog.Infof("connect to mongodb %s successfully on host %s", name, url)
		session.SetMode(mgo.Monotonic, true)
		database = &mgo.Database{
			Name:    name,
			Session: session,
		}
	}
}
