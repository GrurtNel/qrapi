package post

import (
	"gopkg.in/mgo.v2/bson"
	"qrapi/g/x/web"
	"qrapi/x/logger"
	"qrapi/x/mongodb"
	"qrapi/x/validator"
)

var postLog = logger.NewLogger("tbl_post")
var postTable = mongodb.NewTable("post", "post")

type Post struct {
	mongodb.Model `bson:",inline"`
	Title         string `bson:"title" json:"title" validate:"string"`
	Content       string `bson:"content" json:"content" validate:"string,min=0"`
	Description   string `bson:"description" json:"description" validate:"string,min=0"`
	Category      string `bson:"category" json:"category" validate:"string,min=0"`
	Author        string `bson:"author" json:"author"`
	Approve       bool   `bson:"approve" json:"approve"`
}

func (post *Post) Create() error {
	err := validator.Struct(post)
	if err != nil {
		postLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return postTable.Create(post)
}

func (post *Post) Update() error {
	err := validator.Struct(post)
	if err != nil {
		postLog.Error(err)
		return web.WrapBadRequest(err, "")
	}
	return postTable.UpdateId(post.ID, bson.M{
		"$set": bson.M{
			"title":       post.Title,
			"content":     post.Content,
			"category":    post.Category,
			"description": post.Description,
		},
	})
}

func Delete(id string) error {
	return postTable.DeleteByID(id)
}

func Approve(id string) (*Post, error) {
	var post *Post
	var err = postTable.FindId(id).One(&post)
	if err != nil {
		postLog.Error(err)
		return nil, err
	}
	post.Approve = !post.Approve
	return post, postTable.UpdateId(id, post)
}
