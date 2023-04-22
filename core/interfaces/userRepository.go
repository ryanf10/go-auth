package interfaces

import "go-auth/core/entities"

type IUserRepository interface {
	Create(entities.User) entities.User
}
