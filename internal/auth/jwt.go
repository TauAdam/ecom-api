package auth

import (
	"github.com/TauAdam/ecom-api/config"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func CreateJWToken(secret []byte, userID int) (string, error) {
	expirationDuration := time.Second * time.Duration(config.Envs.JWTExpirationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expirationDuration).Unix(),
	})

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
