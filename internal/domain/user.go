package domain

type User struct {
	ID           uint32 `json:"id" db:"id"`
	UUID         string `json:"uuid" db:"uuid"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}
