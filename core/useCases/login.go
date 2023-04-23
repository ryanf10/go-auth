package useCases

import (
	"errors"
	"fmt"
	"go-auth/core/entities"
	"go-auth/core/interfaces"
	error2 "go-auth/core/useCases/error"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Login struct {
	repository interfaces.IUserRepository
}

func NewLogin(repository interfaces.IUserRepository) *Login {
	p := new(Login)
	p.repository = repository
	return p
}

func (login Login) Execute(email string, password string) (entities.User, *string, *error2.RequestError) {
	row := login.repository.FindOneByEmail(email)
	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	switch {
	case err != nil:
		return user, nil, &error2.RequestError{http.StatusForbidden, errors.New("Invalid credentials")}
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errCompare != nil {
		return user, nil, &error2.RequestError{http.StatusForbidden, errors.New("Email and password does not match")}
	}
	token, errToken := CreateToken{}.Execute(user)
	if errToken != nil {
		return user, nil, errToken
	}

	fmt.Println(user)
	return user, &token, nil
}
