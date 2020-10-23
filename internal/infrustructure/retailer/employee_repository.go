package retailer

import (
	"context"
	"database/sql"
	"market-starter/internal/domain/retailer/model"
	interror "market-starter/internal/error"
	"market-starter/internal/error/mysqlerrhandler"
)

type MySQLEmployeeRepository struct {
	db *sql.DB
}




func NewMySQLEmployeeRepository(db *sql.DB) *MySQLEmployeeRepository {
	return &MySQLEmployeeRepository{
		db: db,
	}
}




func (r *MySQLEmployeeRepository) InsertEmployee(ctx context.Context, employee *model.Employee) {
	res, err := r.db.ExecContext(ctx, `
	INSERT INTO
		market_starter.retailer_employee (first_name, second_name, email, employee_password)
	VALUE 
		(?, ?, ?, ?)
`,
	employee.FirstName,
	employee.SecondName,
	employee.Email,
	employee.Password)

	interror.Check(err)

	id, err := res.LastInsertId()
	interror.Check(err)

	employee.ID = int(id)
}

func (r *MySQLEmployeeRepository) GetEmployee(ctx context.Context, employeeID int) *model.Employee {
	employee := &model.Employee{}

	row := r.db.QueryRowContext(ctx, `
	SELECT 
		id, first_name, second_name, employee_password, email
	FROM
		market_starter.retailer_employee e
	WHERE 
		e.id = ?
`, employeeID)

	if err := scanEmployee(row, employee); mysqlerrhandler.NotFound(err) {
		return nil
	}

	employee.Retailers = r.getEmployeeRetailers(ctx, employee)

	return employee
}

func (r *MySQLEmployeeRepository) GetEmployeeByEmail(ctx context.Context, email string) *model.Employee {
	employee := &model.Employee{}

	row := r.db.QueryRowContext(ctx, `
	SELECT 
		id, first_name, second_name, employee_password, email
	FROM
		market_starter.retailer_employee e
	WHERE 
		e.email = ?
`, email)

	if err := scanEmployee(row, employee); mysqlerrhandler.NotFound(err) {
		return nil
	}

	employee.Retailers = r.getEmployeeRetailers(ctx, employee)

	return employee
}

func (r *MySQLEmployeeRepository) InsertRetailerForRetailerEmployee(ctx context.Context, employee *model.Employee, retailer *model.Retailer) {
	tx, err := r.db.BeginTx(ctx, nil)
	interror.Check(err)

	res, err := tx.ExecContext(ctx, `
    INSERT INTO
		market_starter.retailer (name, secret)
	VALUES 
		(?, ?)

`, retailer.Name, retailer.Secret)
	if err != nil {
		tx.Rollback()
		interror.Check(err)
	}

	id, err := res.LastInsertId()
	interror.Check(err)
	retailer.ID = int(id)

	_, err = tx.ExecContext(ctx, `
	INSERT INTO
		market_starter.retailer_employee_role_in_retailer (retailer_id, retailer_employee_id, retailer_employee_role_id)
	VALUES
		(?, ?, ?)
`, retailer.ID, employee.ID, retailer.EmployeeRole.ID)
	if err != nil {
		tx.Rollback()
		interror.Check(err)
	}

	tx.Commit()
}



func (r *MySQLEmployeeRepository) getEmployeeRetailers(ctx context.Context, employee *model.Employee) []*model.Retailer {
	retailers := make([]*model.Retailer, 0)

	rows, err := r.db.QueryContext(ctx, `
	SELECT 
		r.id, r.name, r.secret, rl.id, rl.role
	FROM 
		market_starter.retailer r
	JOIN
		market_starter.retailer_employee_role_in_retailer rer ON rer.retailer_id = r.id AND rer.retailer_employee_id = ?
	JOIN
		market_starter.retailer_employee_role rl ON rl.id = rer.retailer_employee_role_id
`, employee.ID)
	interror.Check(err)
	defer rows.Close()

	for rows.Next() {
		retailer := &model.Retailer{}
		role := &model.RetailerEmployeeRole{}

		err := rows.Scan(&retailer.ID, &retailer.Name, &retailer.Secret, &role.ID, &role.Role)
		interror.Check(err)

		retailer.EmployeeRole = role
		retailers = append(retailers, retailer)
	}
	interror.Check(rows.Err())

	return retailers
}



func scanEmployee(row *sql.Row, employee *model.Employee) error {
	return row.Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.SecondName,
		&employee.Hash,
		&employee.Email,
	)
}

