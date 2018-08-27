package category

import (
	"qrapi/g/x/web"
	"qrapi/o/category"

	"github.com/gin-gonic/gin"
)

type CategoryServer struct {
	*gin.RouterGroup
	web.JsonRender
}

func NewCategoryServer(parent *gin.RouterGroup, name string) *CategoryServer {
	var s = CategoryServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("list", s.getCategories)
	s.POST("create", s.createCategory)
	s.GET("delete", s.deleteCategory)
	return &s
}

func (s *CategoryServer) getCategories(c *gin.Context) {
	categories, err := category.GetCategories()
	web.AssertNil(err)
	s.SendData(c, categories)
}

func (s *CategoryServer) createCategory(c *gin.Context) {
	var cat *category.Category
	c.BindJSON(&cat)
	web.AssertNil(cat.Create())
	s.SendData(c, cat)
}

func (s *CategoryServer) deleteCategory(c *gin.Context) {
	var catID = c.Query("id")
	web.AssertNil(category.DeleteCategoryByID(catID))
	s.Success(c)
}
