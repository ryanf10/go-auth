package useCases

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-auth/core/entities"
	error2 "go-auth/core/useCases/error"
	"net/http"
	"os"
)

type VerifyToken struct{}

func (verifyToken VerifyToken) Execute(tokenString string) (entities.User, *error2.RequestError) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(*TokenClaim); ok && token.Valid {
		return claims.User, nil
	} else {
		fmt.Println(err)
		return claims.User, &error2.RequestError{http.StatusForbidden, errors.New("Invalid token")}
	}

}
