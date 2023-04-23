package useCases

import (
	"database/sql"
	"errors"
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
	var role entities.Role
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &role.ID, &role.Name, &user.CreatedAt)
	user.Role = role

	switch {
	case err == sql.ErrNoRows:
		return user, nil, &error2.RequestError{http.StatusForbidden, errors.New("Invalid credentials")}
	case err != nil:
		return user, nil, &error2.RequestError{http.StatusInternalServerError, errors.New("Something went wrong, please try again later")}
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errCompare != nil {
		return user, nil, &error2.RequestError{http.StatusForbidden, errors.New("Email and password does not match")}
	}
	token, errToken := CreateToken{}.Execute(user)
	if errToken != nil {
		return user, nil, errToken
	}

	return user, &token, nil
}
