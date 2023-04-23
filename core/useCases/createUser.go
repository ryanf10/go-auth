package useCases

import (
	"fmt"
	"github.com/google/uuid"
	"go-auth/core/entities"
	"go-auth/core/interfaces"
	error2 "go-auth/core/useCases/error"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type CreateUser struct {
	userRepository interfaces.IUserRepository
	roleRepository interfaces.IRoleRepository
}

func NewCreateUser(userRepository interfaces.IUserRepository, roleRepository interfaces.IRoleRepository) *CreateUser {
	p := new(CreateUser)
	p.userRepository = userRepository
	p.roleRepository = roleRepository
	return p
}

func (createUser CreateUser) Execute(name string, email string, password string) (entities.User, *error2.RequestError) {
	encyptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	// assing to role 'user'
	var role entities.Role
	row := createUser.roleRepository.FindOneByName("user")
	row.Scan(&role.ID, &role.Name)

	newUser := entities.User{ID: uuid.New().String(), Name: name, Email: email, Password: string(encyptedPassword), RoleId: role.ID, Role: role, CreatedAt: time.Now()}
	user, errorCreate := createUser.userRepository.Create(newUser)
	if errorCreate != nil {
		return user, &error2.RequestError{http.StatusInternalServerError, errorCreate}
	}
	return user, nil
}
