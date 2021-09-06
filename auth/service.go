package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "somethingVerySecret"

type AuthInterface interface {
	Authenticate(reqUser *User) (*User, error)
	GenerateToken(user *User) (*Token, error)
	ParseToken(ctx *gin.Context) (jwt.MapClaims, error)
}

type AuthService struct {
	repoService repoInterface
}

func (authService *AuthService) Authenticate(reqUser *User) (*User, error) {
	dbUser, err := authService.repoService.getUserByUsername(reqUser.Username)
	if err != nil {
		return nil, err
	}
	if reqUser.Username == dbUser.Username && reqUser.Password == dbUser.Password {
		return dbUser, nil
	}

	return nil, errors.New("wrong credentials")
}

func (authService *AuthService) GenerateToken(user *User) (*Token, error) {
	accessToken, err := prepareAccessToken(user)
	if err != nil {
		return nil, err
	}

	// prepare refresh token here, if needed

	token := &Token{
		AccessToken: accessToken,
	}

	return token, nil
}

func prepareAccessToken(user *User) (string, error) {
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

func (authService *AuthService) ParseToken(ctx *gin.Context) (jwt.MapClaims, error) {
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

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	bearToken := ctx.GetHeader("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}

	return "", errors.New("no token provided")
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

func NewAuthService(repo repoInterface) *AuthService {
	return &AuthService{
		repoService: repo,
	}
}
