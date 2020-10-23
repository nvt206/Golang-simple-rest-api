package middlewares

import (
	"demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtService := services.NewJWTService()
		tokenString := jwtService.GetToken(context)
		token,_ := jwtService.ValidateToken(tokenString)
		if !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"err":"Invalid token",
			})
		}
	}

}
