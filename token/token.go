package token

import (
	"crypto/ed25519"
	"crypto/sha512"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ed25519PrivateKey ed25519.PrivateKey
	ed25519PublicKey  ed25519.PublicKey
	secret            = "vvvvvvvbedcjss"
)

func init() {
	seed := make([]byte, ed25519.SeedSize)
	reader := strings.NewReader(secret)
	if _, err := io.ReadFull(reader, seed); err != nil {
		hash := sha512.Sum512([]byte(secret))
		copy(seed, hash[:ed25519.SeedSize])
	}
	privKey := ed25519.NewKeyFromSeed(seed)
	pubKey := privKey.Public().(ed25519.PublicKey)
	//pubKey, privKey, _ := ed25519.GenerateKey(nil)
	ed25519PrivateKey = privKey
	ed25519PublicKey = pubKey
}

type Claims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func GenerateAccessToken(phone string, expirationTime time.Duration) (string, error) {
	expiration := time.Now().Add(expirationTime)
	claims := &Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(ed25519PrivateKey)
}

func GenerateRefreshToken(phone string, expirationTime time.Duration) (string, error) {
	expiration := time.Now().Add(expirationTime)
	claims := &Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(ed25519PrivateKey)
}

func ParseJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return ed25519PublicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
