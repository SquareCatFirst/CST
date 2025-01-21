package CheArticle

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

func Loadarticles(c *gin.Context) {
	id := c.Param("articleid")
	var article Types.Article
	err := DB.Db.Model(&article).Where("id = ?", id).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"title": article.Title, "content": article.Content, "authoruid": article.Author})
}
