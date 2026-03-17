package handlers

import (
	"blog/database"
	"blog/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func HandleLogin(c *gin.Context) {
	session := sessions.Default(c)
	kullaniciadi := c.PostForm("kullaniciadi")
	sifre := c.PostForm("sifre")

	var kullanici models.Kullanici

	if err := database.DB.Where("kullanici_adi=? AND sifre=?", kullaniciadi, sifre).First(&kullanici).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"hata": "Erişim Engellendi: Hatalı Bilgi"})
		return
	}

	session.Set("kullanici_id", kullanici.ID)

	session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
	})

	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"hata": "Oturum hatası oluştu"})
		return
	}

	c.Redirect(http.StatusFound, "/admin/dashboard")
}
