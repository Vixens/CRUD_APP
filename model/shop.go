package model

import (
	"github.com/jinzhu/gorm"
)

type Shop struct {
	gorm.Model
	Shopid      string `json:"shopid"`
	Shopname    string `json:"shopname"`
	Shopaddress string `json:"shopaddress"`
	Shopgenre   string `json:"shopgenre"`
	Shoptel     string `json:"shoptel"`
	Shopsummary string `json:"shopsummary"`
	Shoptable   string `json:"shoptable"`
}

func GetAllShops(db *gorm.DB) ([]Shop, error) {
	shops := []Shop{}
	if err := db.Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func GetShopByID(db *gorm.DB, Shopid uint) (*Shop, error) {
	shop := new(Shop)
	if err := db.Where("Shopid = ?", Shopid).First(&shop).Error; err != nil {
		return nil, err
	}
	return shop, nil
}

func CreateShop(db *gorm.DB, shop *Shop) error {
	return db.Create(&shop).Error
}

func DeleteShopByID(db *gorm.DB, Shopid uint) error {
	return db.Where("Shopid = ?", Shopid).Delete(Shop{}).Error
}
