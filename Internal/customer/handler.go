package customer

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCustomerHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		customerID := r.PathValue("id")

		customer, err := GetCustomer(r.Context(), pool, customerID)
		if err != nil {
			http.Error(w, "customer not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func GetCustomersHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := GetCustomers(r.Context(), pool)

		if err != nil {
			http.Error(w, "failed to retrieve customers", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(customers)
	}
}

func CreateCustomerHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateCustomerRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		customer, err := CreateCustomer(r.Context(), pool, request)

		if err != nil {
			http.Error(w, "failed to create customer", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customer)
	}
}
