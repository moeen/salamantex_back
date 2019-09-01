package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"
)

type Response map[string]interface{}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type Token struct {
	UserID uint
	jwt.StandardClaims
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func GetValidate() *validator.Validate {
	return validate
}
