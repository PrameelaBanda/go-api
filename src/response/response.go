package pkg

type OrdersPayload struct {
OrderName       string `json:"OrderName"`
CustomerCompany string `json:"CustomerCompany"`
CustomerName    string  `json:"CustomerName"`
OrderDate       string  `json:"OrderDate"`
DeliveredAmount float32 `json:"DeliveredAmount"`
TotalAmount     float32 `json:"TotalAmount"`
}