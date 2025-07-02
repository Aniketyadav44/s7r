package utils

import (
	"fmt"
	"time"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/config"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// To generate a new JWT Token
func GenerateToken(userId string) (string, error) {
	if config.JWT_SECRET == "" {
		return "", fmt.Errorf("cannot get jwt secret from env")
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userId,
		"expire_at": time.Now().Add(30 * 24 * time.Hour).Unix(),
	}).SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", fmt.Errorf("cannot create new token: %s", err.Error())
	}

	return token, nil
}

// To validate jwt token and return userId
func ValidateToken(tokenString string) (string, error) {
	if config.JWT_SECRET == "" {
		return "", fmt.Errorf("cannot get jwt secret from env")
	}
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return "", fmt.Errorf("error in generating token for validaton: %s", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error in extracting claims from token")
	}

	if exp, ok := claims["expire_at"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return "", fmt.Errorf("auth token expired")
		}
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("cannot get userId from claims")
	}
	return userId, nil
}

// To create a password hash
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return "", fmt.Errorf("failed to create password hash: %s", err.Error())
	}

	return string(bytes), nil
}

// To check if the hashes match with password
func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
