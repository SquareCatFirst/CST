package GetUser

import (
	"fmt"
	"net/http"
	"spider/DB"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	Types "spider/Type"
)

func GetAUser(c *gin.Context) {
	uid := c.Param("uid")
	var User Types.User
	if _, err := uuid.Parse(uid); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "UID不合法"})
		return
	}
	err := DB.Db.Model(&User).Where("uid = ?", uid).Select()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "查找用户失败"})
	}
	fmt.Println(User)
	c.JSON(http.StatusOK, gin.H{"user": User})
}
