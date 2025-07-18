package models

import "time"

type Employee struct {
	Id         int       `json:"id"`
	EmployeeID string    `json:"employee_id"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	Position   string    `json:"position"`
	Salary     int       `json:"salary"`
	HireDate   time.Time `json:"hire_date"`
	CreatedAt  time.Time `json:"created_at"`
}
