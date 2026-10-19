package api

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	PasswordHash string `json:"password_hash"`
	jwt.RegisteredClaims
}

var tokenKey = newTokenKey()

func newTokenKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic("Не удалось создать ключ подписи: " + err.Error())
	}

	return key
}

func passwordHash(password string) string {
	hash := hmac.New(sha256.New, tokenKey)
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum(nil))
}

func createToken(password string) (string, error) {
	now := time.Now()

	claims := tokenClaims{
		PasswordHash: passwordHash(password),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(8 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(tokenKey)
}

func validToken(value, password string) bool {
	claims := &tokenClaims{}

	token, err := jwt.ParseWithClaims(
		value,
		claims,
		func(_ *jwt.Token) (any, error) {
			return tokenKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
	)
	if err != nil || !token.Valid {
		return false
	}

	return hmac.Equal(
		[]byte(claims.PasswordHash),
		[]byte(passwordHash(password)),
	)
}
