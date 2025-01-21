package CheArticle

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(c *gin.Context) {
	id := c.Param("articleid")
	_, err := DB.Db.Model(&Types.Article{}).Where("id = ?", id).Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "删除文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
