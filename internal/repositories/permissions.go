package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (p *PermissionsProstgres) PermissionsAll(ctx context.Context) ([]*domain.Permission, error) {
	var permissions []*domain.Permission
	err := p.db.SelectContext(ctx, &permissions, "SELECT id, uuid, name from permissions")
	if errors.Is(err, sql.ErrNoRows) {
		return permissions, nil
	}

	if err != nil {
		return nil, fmt.Errorf("cannot get permissions: %w", internal.ErrReadRows)
	}
	return permissions, nil
}

func (p *PermissionsProstgres) PermissionByUUID(ctx context.Context, uuid string) (*domain.Permission, error) {
	var permission = &domain.Permission{}
	query := "SELECT id, uuid, name FROM permissions WHERE uuid=$1"
	err := p.db.GetContext(ctx, permission, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}

	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}
	return permission, nil
}

func (p *PermissionsProstgres) DeletePermission(ctx context.Context, uuid string) error {
	query := "DELETE FROM permissions WHERE uuid=$1"
	_, err := p.db.ExecContext(ctx, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotDelete, uuid)
	}
	return nil
}

func (p *PermissionsProstgres) UpdatePermission(ctx context.Context, permission *domain.Permission) (*domain.Permission, error) {
	query := "UPDATE permissions SET name=$1 WHERE uuid=$2"
	_, err := p.db.ExecContext(ctx, query, permission.Name, permission.UUID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, permission.UUID)
	}
	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, permission.UUID)
	}
	return permission, nil
}
