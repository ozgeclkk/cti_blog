package routes

import (
	"blog/handlers"
	"blog/middleware"
	"net/http" 

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	//Oturum Ayarları
	store := cookie.NewStore([]byte("gizli-anahtar"))
	r.Use(sessions.Sessions("mysession", store))

	
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")

	// Herkese açık
	r.GET("/", handlers.GetPosts)
	r.GET("/login", handlers.ShowLogin)
	r.POST("/login", handlers.HandleLogin)
	r.GET("/post/:id", handlers.GetPostDetail)

	// Çıkış Yapma Rotası
	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear() 
		session.Save()  

		
		c.Redirect(http.StatusFound, "/")
	})
	// Admin paneli
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired)
	{
		
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
