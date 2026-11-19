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
