package account

import (
	"fmt"
	"gihub.com/moeen/salamantex_back/config"
	"gihub.com/moeen/salamantex_back/controllers"
	"gihub.com/moeen/salamantex_back/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type RegisterModel struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

func Register(context *gin.Context) {
	var jsonData RegisterModel
	var errorResponse controllers.ErrorResponse

	if err := context.ShouldBindJSON(&jsonData); err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Body is not valid")
		context.JSON(400, errorResponse)
		return
	}

	if errs := controllers.GetValidate().Struct(jsonData); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errorResponse.Errors = append(errorResponse.Errors, fmt.Sprintf("%v is not valid", err.Field()))
		}

		context.JSON(400, errorResponse)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(jsonData.Password), bcrypt.DefaultCost)
	jsonData.Password = string(hashedPassword)

	user := models.User{
		Name:     jsonData.Name,
		Email:    jsonData.Email,
		Password: jsonData.Password,
	}

	if err := models.GetDB().Create(&user).Error; err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Duplicate data")
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
