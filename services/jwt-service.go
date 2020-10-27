package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTService interface {

	GenerateTokenWithClaims(claims jwt.Claims) string
	ValidateToken(tokenString string) (*jwt.Token,error)
	GetToken(ctx *gin.Context) string
	GetClaims(ctx *gin.Context) *MyCustomClaims
	IsPermit(id uint,ctx *gin.Context) bool
}

type jwtService struct {
	secretKey string
}

func (j jwtService) IsPermit(id uint,ctx *gin.Context) bool {
	return id == j.GetClaims(ctx).ID || j.GetClaims(ctx).IsAdmin
}

func (j jwtService) GetToken(ctx *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	return tokenString
}

func (j jwtService) GetClaims(ctx *gin.Context) *MyCustomClaims {
	tokenString := j.GetToken(ctx)
	token,_ := j.ValidateToken(tokenString)
	if claims,ok := token.Claims.(*MyCustomClaims);ok{
		return claims
	}
	return nil
}

type MyCustomClaims struct {
	ID uint `json:"id"`
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

func (j jwtService) GenerateTokenWithClaims(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	t,err := token.SignedString([]byte(j.secretKey))
	if err!=nil{
		panic(err)
	}
	return t

}



func (j jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {

	return jwt.ParseWithClaims(tokenString,&MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

}

func NewJWTService() JWTService {
	return &jwtService{secretKey: "ACCESS_TOKEN"}

}