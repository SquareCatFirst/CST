package Register

import (
	"fmt"
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
)

type RegRequest struct {
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// api/register
func Register(c *gin.Context) {
	var RegReq RegRequest
	if err := c.ShouldBindJSON(&RegReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
		return
	}
	user := Types.User{
		Uid:      "",
		Nickname: RegReq.Nickname,
		Username: RegReq.Username,
		Phone:    RegReq.Phone,
		Email:    RegReq.Email,
		Password: RegReq.Password,
		Admin:    RegReq.Admin,
		Ban:      false,
	}
	fmt.Println("111")
	var A Types.User
	if err := DB.Db.Model(&A).Where("username = ?", user.Username).Select(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "该用户名已存在"})
		return
	}

	if err := DB.Db.Model(&A).Where("email= ?", user.Email).Select(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "该邮箱已存在"})
		return
	}

	if err := DB.Db.Model(&A).Where("phone = ?", user.Phone).Select(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "该手机号已存在"})
		return
	}
	fmt.Println(user)
	_, err := DB.Db.Model(&user).Insert()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "注册失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
