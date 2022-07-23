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

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
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
