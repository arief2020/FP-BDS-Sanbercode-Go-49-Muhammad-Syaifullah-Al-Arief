package middlewares

import (
	"belajarGin/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        err := token.TokenValid(c)
        if err != nil {
            c.String(http.StatusUnauthorized, err.Error())
            c.Abort()
            return
        }
        c.Next()
    }
}