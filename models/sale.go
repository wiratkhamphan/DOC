package models

import "time"

type Sale struct {
	Id          int       `json:"id"`
	Employee_id string    `json:"employee_id"`
	Customer_id string    `json:"customer_id"`
	Product_id  string    `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
	Sale_date   time.Time `json:"sale_date"`
}
