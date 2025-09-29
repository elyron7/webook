package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/elyron7/webook/internal/web"
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

		clientToken := strings.TrimPrefix(bearerToken, "Bearer ")
		if clientToken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var userClaims web.UserClaims
		token, err := jwt.ParseWithClaims(clientToken, &userClaims, func(token *jwt.Token) (any, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid || userClaims.UserID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		context.Set("userClaims", userClaims)
	}
}
