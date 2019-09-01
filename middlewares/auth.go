package middlewares

import (
	"gihub.com/moeen/salamantex_back/config"
	"gihub.com/moeen/salamantex_back/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthenticationNeeded() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(403, controllers.ErrorResponse{
				Errors: []string{"Token needed"},
			})
			context.Abort()
			return
		}

		slitted := strings.Split(authHeader, " ")
		if len(slitted) != 2 {
			context.JSON(403, controllers.ErrorResponse{
				Errors: []string{"Invalid auth token"},
			})
			context.Abort()
			return
		}

		tokenPart := slitted[1]
		tk := &controllers.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JWTSecret), nil
		})

		if err != nil {
			context.JSON(403, controllers.ErrorResponse{
				Errors: []string{"Malformed authentication token"},
			})
			context.Abort()
			return
		}

		if !token.Valid {
			context.JSON(403, controllers.ErrorResponse{
				Errors: []string{"Token is not valid"},
			})
			context.Abort()
			return
		}

		context.Set("user_id", tk.UserID)
		context.Next()
	}
}
