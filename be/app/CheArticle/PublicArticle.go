package CheArticle

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

func PublicArticle(c *gin.Context) {
	id := c.Param("articleid")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "公开文章失败"})
	}
	_, err := DB.Db.Model(&Types.Article{}).Set("public = ?", true).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "公开文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "公开文章成功"})
}
