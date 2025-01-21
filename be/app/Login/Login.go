package Login

import (
	"fmt"
	"net/http"
	"spider/DB"
	Types "spider/Type"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

// api/login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetLoginUsertype struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Ban      bool   `json:"ban"`
}

func GetLoginUser(db *pg.DB, username string) (R GetLoginUsertype, er error) {
	var user Types.User
	err := db.Model(&user).Where("username = ?", username).Select()
	if err != nil {
		return GetLoginUsertype{
			Uid:      "",
			Username: "",
			Password: "",
		}, err
	}
	return GetLoginUsertype{
		Uid:      user.Uid,
		Username: username,
		Password: user.Password,
		Ban:      user.Ban,
	}, nil

}

func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
		return
	}
	fmt.Print(loginReq)
	uslog, err := GetLoginUser(DB.Db, loginReq.Username)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "查询用户数据失败"})
		fmt.Println("查询用户数据失败")
		return
	}
	if uslog.Ban {
		c.JSON(http.StatusBadRequest, gin.H{"message": "该用户已被封禁！"})
		return
	}
	if uslog.Password == loginReq.Password {
		fmt.Println("登录成功")

		c.SetCookie("uid", uslog.Uid, 3600, "/", "", false, false)

		c.JSON(http.StatusOK, gin.H{})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "用户名或密码错误"})
		fmt.Println("用户名或密码错误")
	}

}
