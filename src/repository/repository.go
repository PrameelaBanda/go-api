package repository

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"go-api/src/models"
	"log"
	"net/http"
	"go-api/src/response"
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

func (r Repository) ListOrders() ([]response.OrdersResponse, error) {
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

	orders := make([]response.OrdersResponse, 0)

	for rows.Next() {
		order := response.OrdersResponse{}
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

func (r Repository) GetOrder(orderName string) (*response.OrdersResponse, error) {
	order := response.OrdersResponse{}
	query := "select o.order_name, cc.company_name, c.name, o.created_at, d.delivered_quantity*oi.price_per_unit as delivered_amount, oi.quantity*oi.price_per_unit as total_amount " +
		"from order_items oi, orders o, deliveries d, customer_companies cc, customers c " +
		"where oi.order_id = o.id and o.customer_id = c.user_id and c.company_id = cc.company_id and d.order_item = oi.order_id " +
		"and o.order_name = $1;"
    fmt.Println("OrderNAme", orderName)
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

func (r Repository) CreateOrders(orders []models.Orders) error {
	var order models.Orders
	insertQry := "insert into orders (id, created_at, order_name, customer_id) values ($1, $2, $3, $4);"
	if len(orders) != 0 {
		for _,order = range orders {
			_, err := r.DB.Exec(insertQry, order.ID, order.CREATED_AT, order.ORDER_NAME, order.CUSTOMER_ID)
			if err != nil {
				return Dberror{
					Code: http.StatusInternalServerError,
					Err:  fmt.Sprintf("error in inserting the data in postgres: %s", err.Error()),
				}
			}
		}
	}
	return nil
}

func (r *Repository) CreateCustomers(customers []models.Customers) error {
	var customer models.Customers
	insertQry := "insert into customers (user_id, login, password, name, company_id, credit_cards) values ($1, $2, $3, $4, $5, $6);"
	if len(customers) != 0 {
		for _,customer = range customers {
			_, err := r.DB.Exec(insertQry, customer.USER_ID, customer.LOGIN, customer.PASSWORD, customer.NAME, customer.COMPANY_ID, pq.Array(customer.CREDIT_CARDS))
			if err != nil {
				return Dberror{
					Code: http.StatusInternalServerError,
					Err:  fmt.Sprintf("error in inserting the data in postgres: %s", err.Error()),
				}
			}
		}
	}
	return nil
}
