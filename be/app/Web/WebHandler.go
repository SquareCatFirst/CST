package Web

import (
	"spider/app/BanUser"
	"spider/app/CheArticle"
	"spider/app/ChePwd"
	"spider/app/GetArticle"
	"spider/app/GetUser"
	"spider/app/LogOut"
	"spider/app/Login"
	"spider/app/Register"
	"spider/app/Search"

	"github.com/gin-gonic/gin"
)

type WebHandler struct{}

func (w *WebHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/login", Login.Login)
	r.POST("/api/register", Register.Register)
	r.POST("/api/articles", CheArticle.SaveArticle)
	r.GET("/api/articles/:articleid", CheArticle.Loadarticles)
	r.PUT("/api/articles/:articleid", CheArticle.SaveArticle)
	r.POST("/api/articles/:articleid/publish", CheArticle.PublicArticle)
	r.DELETE("/api/articles/:articleid", CheArticle.DeleteArticle)
	r.GET("/api/articles/public", GetArticle.GetPublicArticle)
	r.GET("/api/articles/private", GetArticle.GetPrivateArticle)
	r.GET("/api/articles/myarticle/:authoruid", GetArticle.GetMyArticle)
	r.GET("/api/articles/private/:authoruid", GetArticle.GetMyPrivateArticle)
	r.GET("/api/articles/public/:authoruid", GetArticle.GetMyPublicArticle)
	r.GET("/api/search/articles/:context", Search.SearchArticles)
	r.GET("/api/search/users/:context", Search.SearchUsers)
	r.GET("/api/:uid", GetUser.GetAUser)
	r.PUT("/api/users/password/:uid", ChePwd.ChangeUserPwd)
	r.POST("/api/exit/:uid", LogOut.UserLogOut)
	r.POST("/api/ban/:uid", BanUser.BanSomeOne)
	r.POST("/api/unban/:uid", BanUser.UnBanSomeOne)
	r.DELETE("/api/delete/:uid", BanUser.DeleteUser)
	/*	go func() {
		err := r.Run("0.0.0.0:1145")
		if err != nil {
			log.Fatal(err)
		}
	}()*/
}
