package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Role string
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CheckHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}

func GenerateJWT(userId uuid.UUID, role string, secretKey string) (string, error) {
	claims := CustomClaims{
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "mypage",
			Subject:   userId.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	return ss, err
}

func ValidateJWt(tokenString string, secretKey string) (*CustomClaims, error) {
	claimsStruct := &CustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString,
		claimsStruct,
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
		})
	if err != nil {
		return &CustomClaims{}, err
	}
	return claimsStruct, nil
}
func MakeRefreshToken() string {
	key := make([]byte, 32)
	rand.Read(key)
	refreshToken := hex.EncodeToString(key)

	return refreshToken
}
func GetBearerToken(headers http.Header) (string, error) {
	token := strings.TrimSpace(headers.Get("Authorization"))
	if token == "" {
		return "", errors.New("No auth header")
	}
	parts := strings.Fields(token)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" || parts[1] == "" {
		return "", errors.New("Invalid auth header")
	}

	return parts[1], nil
}
