package tokenjwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateJWT(t *testing.T) {
	InitJWTKey("mysecretkey")

	userID := "12345"
	token, err := GenerateJWT(userID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	fmt.Println(token)
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !parsedToken.Valid {
		t.Fatalf("Expected token to be valid")
	}

	if claims.UserID != userID {
		t.Errorf("Expected userID %v, got %v", userID, claims.UserID)
	}

	if claims.ExpiresAt < time.Now().Unix() {
		t.Errorf("Expected token to be valid, but it is expired")
	}
}
