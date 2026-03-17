package routes

import (
	"blog/handlers"
	"blog/middleware"
	"net/http" // Yönlendirmeler için gerekli

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 1. Oturum Ayarları
	store := cookie.NewStore([]byte("gizli-anahtar"))
	r.Use(sessions.Sessions("mysession", store))

	// 2. Dosya Yüklemeleri (Sadece burada kalsın, main.go'dan sil!)
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")

	// --- HERKESE AÇIK ---
	r.GET("/", handlers.GetPosts)
	r.GET("/login", handlers.ShowLogin)
	r.POST("/login", handlers.HandleLogin)
	r.GET("/post/:id", handlers.GetPostDetail)

	// Çıkış Yapma Rotası
	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear() // Tüm session verilerini siler (kullanici_id dahil)
		session.Save()  // Değişikliği kaydet

		// Çıkış yaptıktan sonra ana sayfaya gönder
		c.Redirect(http.StatusFound, "/")
	})
	// --- ADMİN PANELİ (Giriş Şart) ---
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired)
	{
		// Burası artık hem formu hem listeyi gösteren handlers.ShowDashboard'a gidiyor
		admin.GET("/dashboard", handlers.ShowDashboard)

		// Formu gösterme ve kaydetme
		admin.GET("/create", handlers.ShowCreatePost)
		admin.POST("/create", handlers.HandleCreatePost)

		admin.GET("/edit/:id", handlers.ShowEditPost)
		admin.POST("/edit/:id", handlers.HandleUpdatePost)

		// Silme işlemi
		admin.POST("/delete/:id", handlers.DeletePost)
	}
}
