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

func (repo UserRepository) Create(user entities.User) entities.User {
	sql, err := repo.db.Exec("INSERT INTO user (id, name, email, password, roleId, createdAt) VALUES ( ?, ?, ?, ?, 'b0fc999d-5960-4cc9-8549-a5765b959e07', ? )", user.ID, user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sql.LastInsertId())
	return user
}

func (repo UserRepository) FindOneByEmail(email string) *sql.Row {
	row := repo.db.QueryRow("SELECT user.id, user.name, email, password, role.id, role.name, createdAt FROM user, role WHERE user.roleId = role.id AND email = ?", email)
	return row
}
