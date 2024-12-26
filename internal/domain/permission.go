package domain

type Permission struct {
	ID   uint32 `json:"id" db:"id"`
	UUID string `json:"uuid" db:"uuid"`
	Name string `json:"name" db:"name"`
}
