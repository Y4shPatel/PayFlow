package customer

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomer(ctx context.Context, pool *pgxpool.Pool, customerID string) (Customer, error) {
	var customer Customer
	query := `
    SELECT id, full_name, email, phone, created_at
    FROM customers
    WHERE id = $1
`

	err := pool.QueryRow(
		ctx,
		query,
		customerID,
	).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.Email,
		&customer.Phone,
		&customer.CreatedAt)

	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}

func GetCustomers(ctx context.Context, pool *pgxpool.Pool) ([]Customer, error) {
	query := `
		SELECT id, full_name, email, phone, created_at
		FROM customers
	`
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.ID,
			&customer.FullName,
			&customer.Email,
			&customer.Phone,
			&customer.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

func CreateCustomer(
	ctx context.Context,
	pool *pgxpool.Pool,
	request CreateCustomerRequest,
) (Customer, error) {

	var customer Customer

	query := `
		INSERT INTO customers (full_name, email, phone)
		VALUES ($1, $2, $3)
		RETURNING id, full_name, email, phone, created_at
	`

	err := pool.QueryRow(
		ctx,
		query,
		request.FullName,
		request.Email,
		request.Phone,
	).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.Email,
		&customer.Phone,
		&customer.CreatedAt,
	)

	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}
