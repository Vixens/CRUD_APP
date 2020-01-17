package main

import (
	"CRUD_APP/controller"
	"CRUD_APP/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := controller.Server{
		DB: util.InitDB(),
	}
	router := gin.Default()

	router.LoadHTMLGlob("view/*")
	http.HandleFunc("/", moveHandler)
	// router.GET("/", server.GetStorePage)

	/* /loginへのハンドラ(認証プロバイダの選択) */
	http.Handle("/login", &templateHandler{filename: "/login.html"})

	/* 認証用のgomniauthの初期設定 */
	setAuthInfo()

	/* /authへのハンドラ(指定プロバイダでの認証) */
	http.HandleFunc("/auth/", authHandler)

	/* /afterへのハンドラ(認証後のページ・サーブ) */
	http.Handle("/after", &templateHandler{filename: "/after.html"})

	/* /logoutへのハンドラ(認証情報の削除) */
	http.HandleFunc("/logout", logoutHandler)

	router.GET("/new", server.NewStorePage)
	router.GET("/detail/:id", server.GetDetailPage)

	router.POST("/new", server.CreateStoreHandler)
	router.POST("/delete/:id", server.DeleteStoreHandler)
	router.POST("/detail/:id", server.CreateQRHandler)

	router.Run(":8080")
}
