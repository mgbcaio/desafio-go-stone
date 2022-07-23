package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	InvalidTokenErr          = "invalid token"
	TokenParseErr            = "couldn't parse claims"
	TokenExpiredErr          = "token expired"
	InvalidTokenSignatureErr = "token invalid signature"
)

var jwtKey = []byte("super-secret-key")

type JWTClaim struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

type Credentials struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

func GenerateJWT(cpf string, secret string) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Cpf: cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err = errors.New(InvalidTokenSignatureErr)
			return
		}
		return
	}

	if !token.Valid {
		err = errors.New(InvalidTokenErr)
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New(TokenParseErr)
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New(TokenExpiredErr)
		return
	}

	return
}
