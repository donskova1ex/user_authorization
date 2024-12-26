package domain

type Role struct {
	ID   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	UUID string `json:"uuid" db:"uuid"`
}
