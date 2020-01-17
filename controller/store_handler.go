package controller

import (
	"CRUD_APP/model"
	"image/png"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) GetStorePage(c *gin.Context) {
	errMsg := ""
	stores, err := model.GetAllStores(s.DB)
	if err != nil {
		errMsg = "indexページ取得失敗"
		stores = []model.Store{}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"stores": stores,
		"errMsg": &errMsg,
	})
}

func (s *Server) GetDetailPage(c *gin.Context) {
	errMsg := ""
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	store, err := model.GetStoreByID(s.DB, uint(ID))
	if err != nil {
		errMsg = "詳細ページ取得失敗"
		store = &model.Store{}
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"store":  store,
		"errMsg": &errMsg,
	})

}

func (s *Server) CreateQRHandler(c *gin.Context) {

	rand.Seed(time.Now().UnixNano())
	qr_var_random := rand.Float64()
	formated_var_qr := strconv.FormatFloat(qr_var_random, 'f', 4, 64)

	// Create the barcode
	qrCode, _ := qr.Encode(formated_var_qr, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	c.Redirect(http.StatusMovedPermanently, "/detail/:id")

}

func (s *Server) NewStorePage(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{})
}

func (s *Server) CreateStoreHandler(c *gin.Context) {
	c.Request.ParseForm()
	store := new(model.Store)

	store.Storename = c.Request.Form["Storename"][0]
	store.Loc = c.Request.Form["Loc"][0]
	store.Genre = c.Request.Form["Genre"][0]
	store.Tel = c.Request.Form["Tel"][0]
	store.Information = c.Request.Form["Information"][0]
	store.Tables = c.Request.Form["Tables"][0]

	if store.Storename == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "empty:Storename",
		})
		return
	} else if store.Loc == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "empty:Address",
		})
		return
	} else if store.Genre == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "empty:Genre",
		})
		return
	} else if store.Tel == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "empty:Tel",
		})
		return
	} else if store.Information == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "empty:Information",
		})
		return
	} else if store.Tables == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "empty:Tables",
		})
		return
	}
	if err := model.CreateStore(s.DB, store); err != nil {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"errMsg": "登録できませんでした	",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) DeleteStoreHandler(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	if err := model.DeleteStoreByID(s.DB, uint(ID)); err != nil {
		stores, _ := model.GetAllStores(s.DB)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"stores": stores,
			"errMsg": "データ削除に失敗しました。",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
