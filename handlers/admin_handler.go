package handlers

import (
	"blog/database"
	"blog/models"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ShowDashboard(c *gin.Context) {
	var posts []models.Post
	database.DB.Order("olusturma_tarihi desc").Find(&posts)
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"posts": posts})
}

func ShowCreatePost(c *gin.Context) {
	c.HTML(http.StatusOK, "created_post.html", nil)

}

func HandleCreatePost(c *gin.Context) {
	baslik := c.PostForm("title")
	icerik := c.PostForm("content")
	file, err := c.FormFile("image")
	resimYolu := ""

	if err == nil && file != nil {
		if file.Size > 5*1024*1024 {
			c.String(http.StatusBadRequest, "Dosya boyutu çok büyük (Max 5MB)!")
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))
		validExtensions := map[string]bool{
			".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
		}

		if !validExtensions[ext] {
			c.String(http.StatusBadRequest, "Geçersiz dosya formatı! Sadece resim yüklenebilir.")
			return
		}

		yeniDosyaAdi := uuid.New().String() + ext
		resimYolu = "/static/image/" + yeniDosyaAdi

		if err := c.SaveUploadedFile(file, "."+resimYolu); err != nil {
			c.String(http.StatusInternalServerError, "Dosya kaydedilirken bir hata oluştu.")
			return
		}
	}

	yeniPost := models.Post{
		Baslik:          baslik,
		İcerik:          icerik,
		ResimYolu:       resimYolu,
		OlusturmaTarihi: time.Now(),
	}

	database.DB.Create(&yeniPost)
	session := sessions.Default(c)
	session.Save()
	c.Redirect(http.StatusFound, "/admin/dashboard")
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(models.Post{}, id)
	c.Redirect(http.StatusFound, "/admin/dashboard")

}

func ShowEditPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		c.Redirect(http.StatusFound, "/admin/dashboard")
		return
	}

	c.HTML(http.StatusOK, "edit_post.html", gin.H{"post": post})
}

func HandleUpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	database.DB.First(&post, id)

	post.Baslik = c.PostForm("title")
	post.İcerik = c.PostForm("content")

	database.DB.Save(&post)
	c.Redirect(http.StatusFound, "/admin/dashboard")
}
