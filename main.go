package main

import (
	"crud_go/controller"
	"crud_go/util"

	"github.com/gin-gonic/gin"
)

func main() {

	server := controller.Server{
		DB: util.InitDB(),
	}
	router := gin.Default()

	router.LoadHTMLGlob("view/*")

	router.GET("/shops", server.GetShopPage)
	router.GET("/new", server.NewShopPage)
	// router.GET("/shops/:shopid/menu", server.GetMenuHandler)
	// router.GET("/shops/:shopid/menu", server.NewMenuHandler)

	router.POST("/new", server.CreateShopHandler)
	router.POST("/delete/:shopid", server.DeleteShopHandler)

	router.Run(":8080")
}
