package repostiories

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-auth/core/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	p := new(UserRepository)

	p.db = GetDb()
	return p
}

func (repo UserRepository) Create(user entities.User) (entities.User, error) {
	_, err := repo.db.Exec("INSERT INTO user (id, name, email, password, roleId, createdAt) VALUES ( ?, ?, ?, ?, ?, ? )", user.ID, user.Name, user.Email, user.Password, user.RoleId, user.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (repo UserRepository) FindOneByEmail(email string) *sql.Row {
	row := repo.db.QueryRow("SELECT user.id, user.name, email, password, role.id, role.name, createdAt FROM user, role WHERE user.roleId = role.id AND email = ?", email)
	return row
}
