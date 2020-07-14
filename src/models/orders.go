package models

import "time"

type Orders struct {
    ID int `json:"id"`
    CREATED_AT time.Time `json:"created_at"`
    ORDER_NAME string `json:"order_name"`
    CUSTOMER_ID string `json:"customer_id"`
}
