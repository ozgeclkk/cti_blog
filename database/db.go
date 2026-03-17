package database

import (
	"blog/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı")
	}

	db.AutoMigrate(&models.Post{}, &models.Kullanici{})
	DB = db
}

func SeedAdmin() {
	var count int64

	DB.Model(&models.Kullanici{}).Count(&count)

	if count == 0 {
		admin := models.Kullanici{
			KullaniciAdi: "admin",
			Sifre:        "blogadmin",
		}
		DB.Create(&admin)
	}

}
