package repositories

import (
	"context"
	"github.com/donskova1ex/user_authorization/internal"
	"github.com/donskova1ex/user_authorization/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PermissionsProstgres struct {
	db *sqlx.DB
}

func NewPermissionsPostgres(db *sqlx.DB) *PermissionsProstgres {
	return &PermissionsProstgres{db: db}
}

func (p *PermissionsProstgres) CreatePermission(ctx context.Context, permission *domain.Permission) (*domain.Permission, error) {

	var id uint32
	query := "INSERT INTO permissions(uuid, name) VALUES ($1, $2) returning id"
	newUUID := uuid.NewString()
	row := p.db.QueryRowContext(ctx, query, newUUID, permission.Name)
	err := row.Err()
	if err != nil {
		return nil, internal.QueryError
	}

	if err := row.Scan(&id); err != nil {
		return nil, internal.ErrorRowScan
	}
	newPermission := &domain.Permission{
		Name: permission.Name,
		UUID: newUUID,
		ID:   id,
	}
	return newPermission, nil
}
