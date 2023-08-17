package middleware

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type AuhtMiddleware interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type auhtMiddleware struct {

}

var SECRET_KEY = []byte("aiofhioahfahfjjflakfjfljjfljlsj")

func (s *auhtMiddleware) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	
	return signedToken, nil
}

// validate token digunakan agar orang lain tidak bisa membuat jwt token sendri
func (s *auhtMiddleware) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func NewJwtService() *auhtMiddleware{
	return &auhtMiddleware{}
}

