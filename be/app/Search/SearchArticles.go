package Search

import (
	"fmt"
	"net/http"
	"spider/DB"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/html"

	Types "spider/Type"
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
func SearchArticles(c *gin.Context) {
	con := c.Param("context")
	var articles []Types.Article
	if _, err := uuid.Parse(con); err == nil {

		err := DB.Db.Model(&articles).Where("public = ? AND id = ?", true, con).Select()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "查找文章失败"})
		}
		for i := range articles {
			art := &articles[i]
			var user Types.User
			err := DB.Db.Model(&user).Where("uid = ?", art.Author).Select()
			if err != nil {
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
		fmt.Println(articles)
	}
	if len(articles) > 0 {
		c.JSON(http.StatusOK, gin.H{"articles": articles})
		return
	}
	err := DB.Db.Model(&articles).Where("public = ? AND title LIKE ?", true, "%"+con+"%").Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找文章失败"})
	}
	for i := range articles {
		art := &articles[i]
		var user Types.User
		err := DB.Db.Model(&user).Where("uid = ?", art.Author).Select()
		if err != nil {
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
