package models

type Customers struct {
    USER_ID string `json:"USER_ID"`
    LOGIN string `json:"LOGIN"`
    PASSWORD string `json:"PASSWORD"`
    NAME string `json:"NAME"`
    COMPANY_ID int `json:"COMPANY_ID"`
    CREDIT_CARDS []string `json:"CREDIT_CARDS"`
}
