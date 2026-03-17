package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	kullanici := session.Get("kullanici_id")

	if kullanici == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	session.Set("kullanici_id", kullanici)
	session.Save()

	c.Next()

}
