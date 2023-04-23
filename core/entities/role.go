package entities

type Role struct {
	ID   string `db:"id, primarykey" json:"id"`
	Name string `db:"name" json:"name"`
}
