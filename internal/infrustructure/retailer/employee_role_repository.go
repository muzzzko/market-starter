package retailer

import (
	"context"
	"database/sql"
	"market-starter/internal/domain/retailer/model"
	"market-starter/internal/error/mysqlerrhandler"
)

type EmployeeRoleRepository struct {
	db *sql.DB
}



func NewEmployeeRepository(db *sql.DB) *EmployeeRoleRepository {
	return &EmployeeRoleRepository{
		db: db,
	}
}



func (r *EmployeeRoleRepository) GetRoleByName(ctx context.Context, roleName string) *model.RetailerEmployeeRole {
	role := &model.RetailerEmployeeRole{}

	row := r.db.QueryRowContext(ctx, `
	SELECT
		r.id, r.role
	FROM
		market_starter.retailer_employee_role r
	WHERE r.role = ?
`, roleName)

	if err := row.Scan(&role.ID, &role.Role); mysqlerrhandler.NotFound(err) {
		return nil
	}

	return role
}