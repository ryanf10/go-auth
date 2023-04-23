package useCases

import (
	"go-auth/core/entities"
	error2 "go-auth/core/useCases/error"
	"net/http"
	"os"
	"time"
)
import jwt "github.com/golang-jwt/jwt/v5"

type CreateToken struct{}

type TokenClaim struct {
	User entities.User `json:"user"`
	jwt.RegisteredClaims
}

func (createToken CreateToken) Execute(user entities.User) (string, *error2.RequestError) {
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
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return signedToken, &error2.RequestError{http.StatusInternalServerError, err}
	}
	return signedToken, nil
}
