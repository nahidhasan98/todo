package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/todo/model"
)

const secretKey = "somethingVerySecret"

func prepareAccessToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"username":  user.Username,
		"exp":       time.Now().Add(time.Minute * 90).Unix(),
		"tokenType": "access",
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tkn.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateToken(user *model.User) (*model.Token, error) {
	accessToken, err := prepareAccessToken(user)
	if err != nil {
		return nil, err
	}

	// prepare refresh token here, if needed

	token := &model.Token{
		AccessToken: accessToken,
	}

	return token, nil
}

func checkToken(tkn string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tkn, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	bearToken := ctx.GetHeader("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}

	return "", errors.New("no token provided")
}

func ValidateToken(ctx *gin.Context) (jwt.MapClaims, error) {
	token, err := getTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := checkToken(token)
	if err != nil {
		return nil, err
	}

	return *claims, nil
}
