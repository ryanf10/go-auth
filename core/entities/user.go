package entities

import "time"

type User struct {
	ID        string    `db:"id, primarykey" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"createdAt" json:"-"`
	RoleId    string    `db:"roleId" json:"roleId"`

	Role Role `json:"role"`
}
