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
	userRepository interfaces.IUserRepository
	roleRepository interfaces.IRoleRepository
}

func NewCreateUser(userRepository interfaces.IUserRepository, roleRepository interfaces.IRoleRepository) *CreateUser {
	p := new(CreateUser)
	p.userRepository = userRepository
	p.roleRepository = roleRepository
	return p
}

func (createUser CreateUser) Execute(name string, email string, password string) entities.User {
	encyptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	// assing to role 'user'
	var role entities.Role
	row := createUser.roleRepository.FindOneByName("user")
	row.Scan(&role.ID, &role.Name)

	newUser := entities.User{ID: uuid.New().String(), Name: name, Email: email, Password: string(encyptedPassword), RoleId: role.ID, Role: role, CreatedAt: time.Now()}
	return createUser.userRepository.Create(newUser)
}
