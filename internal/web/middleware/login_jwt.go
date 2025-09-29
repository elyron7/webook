package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginJwtMiddlewareBuilder struct {
	paths []string
}

func NewLoginJwtMiddlewareBuilder() *LoginJwtMiddlewareBuilder {
	return &LoginJwtMiddlewareBuilder{}
}

func (b *LoginJwtMiddlewareBuilder) IgnorePaths(paths string) *LoginJwtMiddlewareBuilder {
	b.paths = append(b.paths, paths)
	return b
}

func (b LoginJwtMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		if slices.Contains(b.paths, context.Request.URL.Path) { // ignore paths
			return
		}

		// Use jwt token to verify the user
		bearerToken := context.GetHeader("Authorization")
		if bearerToken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			return []byte("secret"), nil
		})

		if err != nil || !claims.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
