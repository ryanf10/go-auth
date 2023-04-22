package useCases

import (
	"fmt"
	"github.com/google/uuid"
	"go-auth/core/entities"
	"go-auth/core/interfaces"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CreateUser struct {
	repository interfaces.IUserRepository
}

func NewCreateUser(repository interfaces.IUserRepository) *CreateUser {
	p := new(CreateUser)
	p.repository = repository
	return p
}

func (createUser CreateUser) Execute(name string, email string, password string) entities.User {
	encyptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	newUser := entities.User{ID: uuid.New().String(), Name: name, Email: email, Password: string(encyptedPassword), CreatedAt: time.Now()}
	return createUser.repository.Create(newUser)
}
