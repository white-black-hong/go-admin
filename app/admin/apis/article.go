package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/tools/app"
	"net/http"
)

// GetArticleList 获取文章列表
func GetArticleList(c *gin.Context) {

	var res app.Response
	res.Data = "hello world"

	c.JSON(http.StatusOK, res.ReturnOK())
}
