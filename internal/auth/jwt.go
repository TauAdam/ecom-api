package auth

import (
	"fmt"
	"github.com/TauAdam/ecom-api/config"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
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

func JWTGuard(handlerFunc http.HandlerFunc, store models.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := retrieveToken(r)

		token, err := validateToken(tokenStr, []byte(config.Envs.JWTSecret))
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		//	TODO: set the user in the context
	}
}

func retrieveToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return ""
	}
	return bearerToken
}

func validateToken(tokenStr string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}
