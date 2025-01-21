package BanUser

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	uid := c.Param("uid")
	_, err := DB.Db.Model(&Types.User{}).Where("uid = ?", uid).Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "删除用户失败"})
		return
	}
	_, err = DB.Db.Model(&Types.Article{}).Where("author = ?", uid).Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "删除用户成功删除用户文章失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除用户成功"})
}
