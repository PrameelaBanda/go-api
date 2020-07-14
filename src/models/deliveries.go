package models

type Deliveries struct {
   ID                 int  `json:"ID"`
   ORDER_ITEM         int  `json:"ORDER_ITEM"`
   DELIVERED_QUANTITY int  `json:"DELIVERED_QUANTITY"`
}
