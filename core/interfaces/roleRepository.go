package interfaces

import (
	"database/sql"
)

type IRoleRepository interface {
	FindOneByName(name string) *sql.Row
}
