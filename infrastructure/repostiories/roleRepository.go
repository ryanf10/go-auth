package repostiories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository() *RoleRepository {
	p := new(RoleRepository)

	p.db = GetDb()
	return p
}

func (repo RoleRepository) FindOneByName(name string) *sql.Row {
	row := repo.db.QueryRow("SELECT id, name FROM role WHERE name = ?", name)
	return row
}
