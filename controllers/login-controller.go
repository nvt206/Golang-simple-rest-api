package controllers

import (
	"demo/models"
	"demo/services"
	"demo/validations"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

var (
	UserService = services.NewUserService()
)
func Login(ctx *gin.Context) {

	var validation validations.LoginValidation
	if err:= ctx.ShouldBindJSON(&validation);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"errors":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}
	email := validation.Email
	password := validation.Password
	user := UserService.FindOne(&models.User{Email: email})
	if user == nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Incorrect username or password",
		})
		return
	}

	if user.Password != password {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Incorrect username or password",
		})
		return
	}
	claims := services.MyCustomClaims{
		user.ID,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*5).Unix(),
		},
	}

	jwtService := services.NewJWTService()
	token := jwtService.GenerateTokenWithClaims(claims)
	ctx.JSON(http.StatusOK,gin.H{
		"token":token,
	})
}