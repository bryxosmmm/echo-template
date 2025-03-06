package tokenjwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey string

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	fmt.Println("Generated Token:", signedToken)
	return signedToken, nil
}

func DecodeJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func InitJWTKey(key string) {
	SecretKey = key
	fmt.Println("✅ JWT key initialized successfully")
}
