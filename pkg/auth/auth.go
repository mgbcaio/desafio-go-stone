package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
)

const (
	InvalidTokenErr          = "invalid token"
	TokenParseErr            = "couldn't parse claims"
	TokenExpiredErr          = "token expired"
	InvalidTokenSignatureErr = "token invalid signature"
	UnauthorizedErr          = "unauthorized"
	BadRequestErr            = "bad request"
)

var jwtKey = []byte("super-secret-key")

type Claims struct {
	Cpf       string `json:"cpf"`
	UserID    int64  `json:"user_id"`
	AccountID int64  `json:"account_id"`
	jwt.StandardClaims
}

type Credentials struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

func getAccountAndUserIDs(cpf string) (int64, int64) {
	var (
		userID    int64
		accountID int64
	)

	for _, user := range mocks.Users {
		if user.Cpf == cpf {
			userID = user.Id
		}
	}

	for _, account := range mocks.Accounts {
		if account.Cpf == cpf {
			accountID = account.Id
		}
	}
	return userID, accountID
}

// GenerateJWT generates a JWT token attaching some claims to the token. Returns the token string, the expiration time and an error if something goes wrong.
func GenerateJWT(cpf string) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	userID, accountID := getAccountAndUserIDs(cpf)

	claims := &Claims{
		Cpf:       cpf,
		UserID:    userID,
		AccountID: accountID,
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

// ValidateToken validates the given token string. returns an error if something goes wrong.
func ValidateToken(signedToken string) (err error) {
	token, _, err := ParseToken(signedToken)
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

	claims, ok := token.Claims.(*Claims)
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

// ExtractAndValidateToken extracts and validates the token. returns an error if something goes wrong.
func ExtractAndValidateToken(r *http.Request) (err error) {
	token, err := ExtractToken(r)
	if err != nil {
		return
	}

	err = ValidateToken(token)
	if err != nil {
		if strings.Contains(err.Error(), InvalidTokenSignatureErr) || strings.Contains(err.Error(), TokenExpiredErr) || strings.Contains(err.Error(), InvalidTokenErr) {
			err = errors.New(UnauthorizedErr)
			return
		}
		err = errors.New(BadRequestErr)
		return
	}
	return
}

// ExtractToken extracts the token from the request cookies. Returns an error if something goes wrong.
func ExtractToken(r *http.Request) (token string, err error) {
	cookies, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			err = errors.New(UnauthorizedErr)
			return
		}
		err = errors.New(BadRequestErr)
		return
	}

	token = cookies.Value
	return
}

// ParseToken parses the token string and returns the JWT token, the Claims and an error if something goes wrong.
func ParseToken(signedToken string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}
