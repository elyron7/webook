package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (b *LoginMiddlewareBuilder) IgnorePaths(paths string) *LoginMiddlewareBuilder {
	b.paths = append(b.paths, paths)
	return b
}

func (b LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		if slices.Contains(b.paths, context.Request.URL.Path) { // ignore paths
			return
		}

		session := sessions.Default(context)
		id := session.Get("userId")
		if id == nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updatedAt := session.Get("updatedAt")
		now := time.Now().UnixMilli()
		if updatedAt == nil {
			session.Set("updatedAt", now)
			fmt.Printf("first request, updatedAt: %d\n", now)
			fmt.Printf("userId: %s\n", session.Get("userId"))
			session.Save()
			return
		}

		updatedAtInt, ok := updatedAt.(int64)
		if !ok {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "session updatedAt is not int64"})
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if now-updatedAtInt > 10000 { // 10 seconds
			session.Set("updatedAt", now)
			fmt.Printf("session expired, updatedAt: %d\n", now)
			fmt.Printf("userId: %s\n", session.Get("userId"))
			session.Save()
			return
		}
	}
}
