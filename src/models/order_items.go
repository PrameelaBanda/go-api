package models

type OrderItems struct {
    ID         int  `json:"ID"`
    ORDER_ID   int `json:"ORDER_ID"`
    PRICE_PER_UNIT        float32  `json:"PRICE_PER_UNIT"`
    QUANTITY   int  `json:"QUANTITY"`
    PRODUCT    string  `json:"PRODUCT"`
}
