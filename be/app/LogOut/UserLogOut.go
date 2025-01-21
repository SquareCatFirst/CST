package LogOut

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogOut(c *gin.Context) {
	uid := c.Param("uid")
	c.SetCookie("uid", uid, -1, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{})
}
