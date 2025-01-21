package ChePwd

import (
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

type PwdReq struct {
	Password string `json:"password"`
}

func ChangeUserPwd(c *gin.Context) {
	uid := c.Param("uid")
	var pwd PwdReq
	if err := c.ShouldBindJSON(&pwd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
		return
	}
	_, err := DB.Db.Model(&Types.User{}).Set("password = ?", pwd.Password).Where("uid = ?", uid).Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "修改密码失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "修改密码成功"})
}
