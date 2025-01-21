package BanUser

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

func BanSomeOne(c *gin.Context) {
	uid := c.Param("uid")
	_, err := DB.Db.Model(&Types.User{}).Set("ban = ?", true).Where("uid = ?", uid).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "封禁用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "封禁用户成功"})
}
