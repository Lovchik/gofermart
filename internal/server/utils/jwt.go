package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	log "gitlab.indev.by/pkg/go/logger"
	"gofermart/internal/server/models"
	"strconv"
	"time"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func IsValidToken(signedToken, tokenType string) bool {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return false, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return models.GetJwtPair().PublicKey, nil
	})

	if err != nil {
		log.Error("JWT parse", "err", err.Error())
		return false
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["type"] != tokenType {
		return false
	}
	return true
}

func GenerateJWT(userId int64) (TokenPair, error) {
	var tokenPair TokenPair
	accessToken := jwt.New(jwt.SigningMethodES256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = userId
	claims["type"] = "access"
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(30)).Unix()
	signedToken, err := accessToken.SignedString(models.GetJwtPair().PrivateKey)
	if err != nil {
		return tokenPair, err
	}

	refreshToken := jwt.New(jwt.SigningMethodES256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["user_id"] = userId
	refreshClaims["iat"] = time.Now().Unix()
	refreshClaims["exp"] = time.Now().Add(time.Minute * time.Duration(120)).Unix()
	refreshClaims["type"] = "refresh"
	signedRefreshToken, err := refreshToken.SignedString(models.GetJwtPair().PrivateKey)
	if err != nil {
		return tokenPair, err
	}
	tokenPair.AccessToken = signedToken
	tokenPair.RefreshToken = signedRefreshToken
	return tokenPair, nil

}

func GetUserId(tokenString string) (int64, error) {
	var name string
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		name = fmt.Sprint(claims["user_id"])
	}
	if name == "" {
		return 0, fmt.Errorf("Id не найден")
	}
	userId, err := strconv.ParseInt(name, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
