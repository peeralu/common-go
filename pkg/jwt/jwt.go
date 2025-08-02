package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewClaims(id int32, expire_hour int16, iss string) jwt.MapClaims {
	expire := time.Now().Add(time.Duration(expire_hour) * time.Hour)
	return jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
		"iss": iss,
		"exp": expire.Unix(),
	}
}

func GenerateJWT(secret string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, errors.New("invalid token claims")
}
