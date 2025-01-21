package main

import (
	"spider/DB"
	"spider/app/Web"

	"github.com/gin-gonic/gin"
)

func main() {
	DB.Db = DB.DBinit()
	defer DB.Db.Close()
	router := gin.Default()
	var w Web.WebHandler
	w.RegisterRoutes(router)

	router.Run("0.0.0.0:1145")
}
