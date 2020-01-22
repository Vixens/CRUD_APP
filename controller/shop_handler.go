package controller

import (
	"net/http"
	"strconv"

	// "image/png"
	"CRUD_APP/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/boombuler/barcode"
	// "github.com/boombuler/barcode/qr"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) GetShopPage(c *gin.Context) {
	shops, err := model.GetAllShops(s.DB)
	if err != nil {
		shops = []model.Shop{}
	}
	c.JSON(http.StatusOK, shops)
}

func (s *Server) NewShopPage(c *gin.Context) {
	c.HTML(http.StatusOK, "new.tmpl", gin.H{})
}

func (s *Server) CreateShopHandler(c *gin.Context) {
	c.Request.ParseForm()
	shop := new(model.Shop)

	shop.Shopname = c.Request.Form["Shopname"][0]
	shop.Shopaddress = c.Request.Form["Shopaddress"][0]
	shop.Shopgenre = c.Request.Form["Shopgenre"][0]
	shop.Shoptel = c.Request.Form["Shoptel"][0]
	shop.Shopsummary = c.Request.Form["Shopsummary"][0]
	shop.Shoptable = c.Request.Form["Shoptables"][0]

	if shop.Shopname == "" {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "empty:Shopname",
		})
		return
	} else if shop.Shopaddress == "" {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "empty:Shopaddress",
		})
		return
	} else if shop.Shopgenre == "" {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "empty:Shopgenre",
		})
		return
	} else if shop.Shoptel == "" {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "empty:Shoptel",
		})
		return
	} else if shop.Shopsummary == "" {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "empty:Shopsummary",
		})
		return
	} else if shop.Shoptable == "" {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "empty:Shoptable",
		})
		return
	}
	if err := model.CreateShop(s.DB, shop); err != nil {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{
			"errMsg": "登録できませんでした	",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) DeleteShopHandler(c *gin.Context) {
	ShopID, err := strconv.Atoi(c.Param("Shopid"))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	if err := model.DeleteShopByID(s.DB, uint(ShopID)); err != nil {
		shops, _ := model.GetAllShops(s.DB)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"shops":  shops,
			"errMsg": "データ削除に失敗しました。",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
