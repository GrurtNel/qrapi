package post

import (
	"context"
	"github.com/gin-gonic/gin"
	"qrapi/g/x/math"
	"qrapi/g/x/web"
	"qrapi/o/post"
	// "qrapi/o/push_token"
	// "qrapi/x/fcm"
)

type PostServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewPostServer(parent *gin.RouterGroup, name string) *PostServer {
	var s = PostServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("list", s.getPosts)
	s.POST("create", s.createPost)
	s.POST("update", s.updatePost)
	s.GET("detail/:id", s.getDetail)
	s.GET("delete", s.deletePost)
	s.GET("approve", s.approvePost)
	s.POST("upload/thumb", s.UploadThumb)
	s.POST("upload/post-image", s.UploadPostImage)
	return &s
}

func (s *PostServer) getPosts(c *gin.Context) {
	var userID = c.Query("user_id")
	var posts, err = post.GetPosts(userID)
	web.AssertNil(err)
	s.SendData(c, posts)
}

func (s *PostServer) createPost(c *gin.Context) {
	var post *post.Post
	c.BindJSON(&post)
	web.AssertNil(post.Create())
	s.SendData(c, post)
}

func sendNotiToAllUser(c context.Context) {
	// var pushes, err = push_token.GetAllPushToken()
	// if pushes != nil {
	// 	fcm.Send
	// }
}

func (s *PostServer) updatePost(c *gin.Context) {
	var post *post.Post
	c.BindJSON(&post)
	web.AssertNil(post.Update())
	s.SendData(c, post)
}

func (s *PostServer) getDetail(c *gin.Context) {
	var postID = c.Param("id")
	var post, err = post.GetPost(postID)
	web.AssertNil(err)
	s.SendData(c, post)
}

func (s *PostServer) deletePost(c *gin.Context) {
	var postID = c.Query("id")
	web.AssertNil(post.Delete(postID))
	s.Success(c)
}

func (s *PostServer) approvePost(c *gin.Context) {
	var postID = c.Query("id")
	var post, err = post.Approve(postID)
	web.AssertNil(err)
	s.SendData(c, post)
}

func (s *PostServer) UploadThumb(c *gin.Context) {
	var postID = c.Query("post_id")
	var file, err = c.FormFile("thumb")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/post/"+postID))
	s.Success(c)
}

func (s *PostServer) UploadPostImage(c *gin.Context) {
	var fileName = math.RandString("img", 11)
	var file, err = c.FormFile("post-image")
	web.AssertNil(err)
	web.AssertNil(c.SaveUploadedFile(file, "./upload/post/"+fileName))
	// s.SendData(c, map[string]interface{}{
	// 	"image_url": ORIGIN + c.Request.Host + "/static/post/" + fileName,
	// })
	c.JSON(200, map[string]interface{}{
		"location": ORIGIN + c.Request.Host + "/static/post/" + fileName,
	})
}

const (
	ORIGIN = "http://"
)
