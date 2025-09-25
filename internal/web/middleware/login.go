package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct{}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (b LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		//if session == nil {
		//	context.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}

		if context.Request.URL.Path == "/users/signup" ||
			context.Request.URL.Path == "/users/login" {
			return
		}

		id := session.Get("userId")
		if id == nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			//context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
