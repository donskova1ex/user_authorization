package domain

type Role struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	UUID string `json:"uuid" db:"uuid"`
}
