package BanUser

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

func UnBanSomeOne(c *gin.Context) {
	uid := c.Param("uid")
	_, err := DB.Db.Model(&Types.User{}).Set("ban = ?", false).Where("uid = ?", uid).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "解封用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "解封用户成功"})

}
