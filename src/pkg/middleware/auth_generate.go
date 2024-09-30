package middleware

import (
	"fmt"
	"hisaab-kitaab/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var secretKey = []byte("secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("unexpected signing method:xyz")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		if err.Error() == "token is malformed: could not base64 decode signature: illegal base64 data at input byte 40" {
			return fmt.Errorf("invalid token format: %v", err)
		}
		return fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return fmt.Errorf("could not parse claims")
	}
	logger.Log().Debug("claims", zap.Any("claims", claims))
	return nil
}
