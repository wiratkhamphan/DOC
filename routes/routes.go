package routes

import (
	"dog/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/report_sale", controllers.SalesGet)
	app.Get("/report_sale/:in_id", controllers.GET_sale_by_id)

}
