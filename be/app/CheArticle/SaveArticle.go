package CheArticle

import (
	"fmt"
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

type ArtReq struct {
	ArticleUid string `json:"articleuid"`
	AuthorUid  string `json:"authoruid"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

func SaveArticle(c *gin.Context) {
	var Article ArtReq
	if err := c.ShouldBindJSON(&Article); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
		return
	}
	PutArticle := Types.Article{
		Id:      Article.ArticleUid,
		Author:  Article.AuthorUid,
		Title:   Article.Title,
		Content: Article.Content,
		Public:  false,
	}
	if PutArticle.Title == "" || PutArticle.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "不可保存空文章"})
	}
	if PutArticle.Id == "" {
		_, err := DB.Db.Model(&PutArticle).Insert()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "保存失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "保存成功", "uuid": PutArticle.Id})
	} else {
		_, err := DB.Db.Model(&PutArticle).Set("title = ?", PutArticle.Title).Set("content = ?", PutArticle.Content).Where("id = ?", PutArticle.Id).Update()
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "保存失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "保存成功", "uuid": PutArticle.Id})
	}

}
