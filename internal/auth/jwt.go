package auth

import (
	"context"
	"fmt"
	"github.com/TauAdam/ecom-api/config"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/TauAdam/ecom-api/shared/response"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

const ContextKeyUserID = "userID"

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
			log.Printf("Error validating token: %v", err)
			response.SendError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
			return
		}
		if !token.Valid {
			log.Printf("invalid token: %v", err)
			response.SendError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userID, err := strconv.Atoi(claims["userID"].(string))
		if err != nil {
			log.Printf("Error parsing userID: %v", err)
			response.SendError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
			return
		}
		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("Error getting user: %v", err)
			response.SendError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKeyUserID, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
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

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(ContextKeyUserID).(int)
	if !ok {
		return -1
	}
	return userID
}
