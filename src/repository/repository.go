package repository

import (
	"database/sql"
	"fmt"
	"go-api/src/models"
	"log"
	"net/http"
)

type Repository struct {
	DB         *sql.DB
	DBUri      string
	DBName     string
	DBUsername string
	DBPassword string
	Table      string
}

type Dberror struct {
	Code int
	Err  string
}

func (e Dberror) Error() string {
	return e.Err
}

func (r *Repository) Init() error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", r.DBUsername, r.DBPassword, r.DBUri, r.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	log.Print("db client initialized")
	r.DB = db
	return nil
}

func (r Repository) CloseConn() error {
	return r.DB.Close()
}

func (r Repository) ListOrders() ([]models.OrdersResponse, error) {
	var query = "select o.order_name, cc.company_name, c.name, o.created_at, d.delivered_quantity*oi.price_per_unit as delivered_amount, oi.quantity*oi.price_per_unit as total_amount " +
		"from order_items oi, orders o, deliveries d, customer_companies cc, customers c " +
		"where oi.order_id = o.id and o.customer_id = c.user_id and c.company_id = cc.company_id and d.order_item = oi.order_id;"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, Dberror{
			Code: http.StatusInternalServerError,
			Err:  err.Error(),
		}
	}

	orders := make([]models.OrdersResponse, 0)

	for rows.Next() {
		order := models.OrdersResponse{}
		err := rows.Scan(&order.OrderName, &order.CustomerCompany, &order.CustomerName, &order.OrderDate, &order.DeliveredAmount, &order.TotalAmount)
		if err != nil {
			return nil, Dberror{
				Code: http.StatusInternalServerError,
				Err:  fmt.Sprintf("error in conversion from db row: %s", err.Error()),
			}
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r Repository) GetOrder(orderName string) (*models.OrdersResponse, error) {
	order := models.OrdersResponse{}
	var query = "select o.order_name, cc.company_name, c.name, o.created_at, d.delivered_quantity*oi.price_per_unit as delivered_amount, oi.quantity*oi.price_per_unit as total_amount " +
		"from order_items oi, orders o, deliveries d, customer_companies cc, customers c " +
		"where oi.order_id = o.id and o.customer_id = c.user_id and c.company_id = cc.company_id and d.order_item = oi.order_id" +
		"and o.order_name = $1;"

	row := r.DB.QueryRow(query, orderName)
	err := row.Scan(&order.OrderName, &order.CustomerCompany, &order.CustomerName, &order.OrderDate, &order.DeliveredAmount, &order.TotalAmount)
	switch {
	case err == sql.ErrNoRows:
		return nil, Dberror{
			Code: http.StatusNotFound,
			Err:  fmt.Sprintf("no order found: %s", err.Error()),
		}
	case err != nil:
		return nil, Dberror{
			Code: http.StatusInternalServerError,
			Err:  err.Error(),
		}
	}

	return &order, err
}