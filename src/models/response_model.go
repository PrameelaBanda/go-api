package models

type OrdersResponse struct {
OrderName       string
CustomerCompany string
CustomerName    string
OrderDate       string
DeliveredAmount float32
TotalAmount     float32
}