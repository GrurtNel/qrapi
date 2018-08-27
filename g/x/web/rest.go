package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonRender struct {
}

func (r *JsonRender) SendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":   data,
		"status": "ok",
	})
}
func (r *JsonRender) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":   nil,
		"status": "ok",
	})
}
