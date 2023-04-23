package useCases

import (
	"fmt"
	"go-auth/core/entities"
	error2 "go-auth/core/useCases/error"
	"net/http"
	"time"
)
import jwt "github.com/golang-jwt/jwt/v5"

type CreateToken struct{}

type TokenClaim struct {
	User entities.User `json:"user"`
	jwt.RegisteredClaims
}

func (createToken CreateToken) Execute(user entities.User) (string, *error2.RequestError) {
	fmt.Println(user)
	claims := TokenClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		User: user,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return signedToken, &error2.RequestError{http.StatusInternalServerError, err}
	}
	return signedToken, nil
}
