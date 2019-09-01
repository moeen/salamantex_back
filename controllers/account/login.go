package account

import (
	"gihub.com/moeen/salamantex_back/config"
	"gihub.com/moeen/salamantex_back/controllers"
	"gihub.com/moeen/salamantex_back/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {
	var jsonData LoginModel
	var errorResponse controllers.ErrorResponse

	if err := context.ShouldBindJSON(&jsonData); err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Body is not valid")
		context.JSON(400, errorResponse)
		return
	}

	var user models.User
	if err := models.GetDB().Where(&models.User{Email: jsonData.Email}).Find(&user).Error; err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Email or password may be incorrect.")
		context.JSON(400, errorResponse)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonData.Password)); err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Email or password may be incorrect.")
		context.JSON(400, errorResponse)
		return
	}

	tk := &controllers.Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(config.GetConfig().JWTSecret))

	context.JSON(200, controllers.Response{
		"token": tokenString,
	})
	return
}
