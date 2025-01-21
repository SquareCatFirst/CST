package GetArticle

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func GetMyPublicArticle(c *gin.Context) {
	uid := c.Param("authoruid")
	var articles []Types.Article
	err := DB.Db.Model(&articles).Where("author = ? AND public = ?", uid, true).Order("createtime DESC").Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找文章失败"})
		return
	}
	var user Types.User
	err = DB.Db.Model(&user).Where("uid = ?", uid).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找用户失败"})
		return
	}
	for i := range articles {
		art := &articles[i]
		art.Author = user.Nickname + "   UID:" + art.Author
		doc, err := html.Parse(strings.NewReader(art.Content))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "解析html失败"})
			return
		}
		// 提取文本
		text := extractText(doc)
		if len(text) > 400 {
			art.Content = text[:400]
		}
	}
	c.JSON(http.StatusOK, gin.H{"articles": articles, "authornickname": user.Nickname})
}
