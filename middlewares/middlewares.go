package middlewares

import (
	"final-project/utils/token"
	"fmt"
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
func JwtAdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        value, err := token.TokenAdmin(c)
        fmt.Println(value)
        if err != nil {
            c.String(http.StatusUnauthorized, err.Error())
            c.Abort()
            return
        }
        if value != "admin" {
            c.String(http.StatusUnauthorized, "hanya admin yang boleh !!!")
            c.Abort()
            return
        }
        c.Next()
    }
}