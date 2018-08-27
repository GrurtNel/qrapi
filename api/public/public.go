package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qrapi/g/x/web"
	"qrapi/o/category"
	"qrapi/o/post"
	"qrapi/o/push_token"
	"qrapi/x/fcm"
	"strconv"
)

type PublicServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewPublicServer(parent *gin.RouterGroup, name string) *PublicServer {
	var s = PublicServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("post/list", s.getPosts)
	s.GET("post/detail/:id", s.getDetail)
	s.GET("category/list", s.getCategories)
	s.GET("test", s.test)
	s.POST("push_token/create", s.createPush)
	return &s
}

func (s *PublicServer) test(c *gin.Context) {
	var token = c.Query("token")
	var err, str = fcm.SendToOne(token, fcm.FmcMessage{Title: "Hello", Body: "Anh"})
	fmt.Println(err)
	fmt.Println(str)
	s.Success(c)
}

func (s *PublicServer) getPosts(c *gin.Context) {
	var cateID = c.Query("cat_id")
	var page, _ = strconv.ParseInt(c.Query("page"), 10, 32)
	var posts, err = post.GetAllPosts(cateID, page)
	web.AssertNil(err)
	s.SendData(c, posts)
}

func (s *PublicServer) getCategories(c *gin.Context) {
	var cats, err = category.GetCategories()
	web.AssertNil(err)
	c.JSON(200, cats)
	// s.SendData(c, cats)
}

func (s *PublicServer) getDetail(c *gin.Context) {
	var postID = c.Param("id")
	var post, err = post.GetPost(postID)
	web.AssertNil(err)
	s.SendData(c, post)
}

func (s *PublicServer) createPush(c *gin.Context) {
	var push *push_token.PushToken
	c.BindJSON(&push)
	web.AssertNil(push.Create())
	s.Success(c)
}
