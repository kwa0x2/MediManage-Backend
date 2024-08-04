package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionUserUUID := session.Get("uuid")
		sessionUserRole := session.Get("role")

		if sessionUserUUID == nil || sessionUserRole == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization failed",
			})
			ctx.Abort()
			return
		}

		return
	}
}
