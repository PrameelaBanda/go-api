package models

type OrderItems struct {
    ID         int
    ORDER_ID   int
    PRICE_PER_UNIT        float32
    QUANTITY   int
    PRODUCT    string
}