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

type RolesPostgres struct {
	db *sqlx.DB
}

func NewRolesPostgres(db *sqlx.DB) *RolesPostgres {
	return &RolesPostgres{db: db}
}

func (r *RolesPostgres) CreateRole(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	var id uint32

	query := "INSERT INTO roles(uuid, name) VALUES ($1, $2) RETURNING id"
	newUUID := uuid.NewString()
	row := r.db.QueryRowContext(ctx, query, newUUID, role.Name)
	err := row.Err()

	if err != nil {
		return nil, internal.QueryError
	}
	if err := row.Scan(&id); err != nil {
		return nil, internal.ErrorRowScan
	}
	newRole := &domain.Role{
		Name: role.Name,
		UUID: newUUID,
		ID:   id,
	}
	return newRole, nil
}

func (r *RolesPostgres) RoleAll(ctx context.Context) ([]*domain.Role, error) {
	var roles []*domain.Role
	err := r.db.SelectContext(ctx, &roles, "SELECT id, uuid, name FROM roles")
	if errors.Is(err, sql.ErrNoRows) {
		return roles, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can not get roles: %w", internal.ErrReadRows)
	}
	return roles, nil
}

func (r *RolesPostgres) RoleByUUID(ctx context.Context, uuid string) (*domain.Role, error) {
	var role = &domain.Role{}
	query := "SELECT id, uuid, name FROM roles WHERE uuid=$1"
	err := r.db.GetContext(ctx, role, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}
	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}
	return role, nil
}

func (r *RolesPostgres) DeleteRole(ctx context.Context, uuid string) error {
	query := "DELETE FROM roles WHERE uuid=$1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}
	return nil
}

func (r *RolesPostgres) UpdateRole(ctx context.Context, role *domain.Role) error {
	query := "UPDATE roles SET name=$1 WHERE uuid=$2"
	_, err := r.db.ExecContext(ctx, query, role.Name, role.UUID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, role.UUID)
	}
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, role.UUID)
	}
	return nil
}
