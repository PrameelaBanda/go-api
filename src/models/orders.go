package models

import "time"

type Orders struct {
    ID int
    CREATED_AT time.Time
    ORDER_NAME string
    CUSTOMER_ID string
}