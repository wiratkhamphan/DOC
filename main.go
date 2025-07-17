package main

import (
	"context"
	"dog/condb"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Sale struct {
	Id          int       `json:"id"`
	Employee_id string    `json:"employee_id"`
	Customer_id string    `json:"customer_id"`
	Product_id  string    `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
	Sale_date   time.Time `json:"sale_date"`
}

func main() {
	condb.DB_Lek()

	app := fiber.New()
	app.Get("/", GET_sales)
	app.Listen(":8080")
}

func GET_sales(c *fiber.Ctx) error {
	ctx := context.Background()
	conn, err := condb.DB_Lek()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, "SELECT id, employee_id, customer_id, product_id, quantity, total_price, sale_date FROM sales")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	var sales []Sale
	for rows.Next() {
		var s Sale
		err := rows.Scan(&s.Id, &s.Employee_id, &s.Customer_id, &s.Product_id, &s.Quantity, &s.TotalPrice, &s.Sale_date)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		sales = append(sales, s)
	}

	return c.JSON(sales)
}
