package handlers

import (
	"blog/database"
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post

	database.DB.Order("olusturma_tarihi desc").Find(&posts)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": posts,
	})

}

func GetPostDetail(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// Veritabanından veriyi çek
	if err := database.DB.First(&post, id).Error; err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// templates/blog/post_detail.html dosyasını render et
	c.HTML(http.StatusOK, "post_detail.html", gin.H{
		"post": post,
	})
}
