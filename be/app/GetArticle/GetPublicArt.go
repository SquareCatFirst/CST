package GetArticle

import (
	"fmt"
	"net/http"
	"spider/DB"
	Types "spider/Type"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}

	return text
}
func GetPublicArticle(c *gin.Context) {
	var articles []Types.Article
	err := DB.Db.Model(&articles).Where("public = ?", true).Order("createtime DESC").Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找文章失败"})
		return
	}
	for i := range articles {
		art := &articles[i]
		var user Types.User
		err := DB.Db.Model(&user).Where("uid = ?", art.Author).Select()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "查找用户失败"})
			return
		}
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
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}
