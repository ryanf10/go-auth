package interfaces

import (
	"database/sql"
	"go-auth/core/entities"
)

type IUserRepository interface {
	Create(user entities.User) (entities.User, error)
	FindOneByEmail(email string) *sql.Row
}
