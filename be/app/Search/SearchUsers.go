package Search

import (
	"net/http"
	"spider/DB"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	Types "spider/Type"
)

// 支持昵称模糊匹配
// uid 用户名 邮箱 手机号精确匹配 管理员可以搜索邮箱 手机号
func SearchUsers(c *gin.Context) {
	con := c.Param("context")
	var Users []Types.User
	if _, err := uuid.Parse(con); err == nil {

		err := DB.Db.Model(&Users).Where("uid = ?", con).Select()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "查找用户失败"})
		}

	}
	if len(Users) > 0 {
		c.JSON(http.StatusOK, gin.H{"users": Users})
		return
	}
	err := DB.Db.Model(&Users).Where("nickname LIKE ? OR username = ? ", "%"+con+"%", con).Select()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找用户失败"})
	}
	c.JSON(http.StatusOK, gin.H{"users": Users})
}
