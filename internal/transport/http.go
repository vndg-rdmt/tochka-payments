package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vndg-rdmt/paymentspod/internal/controller"
)

func NewHttp(ctr *controller.Fiber, host string) error {
	app := fiber.New()
	app.Post("/api/payment", ctr.Payment)
	app.Get("/api/status/:id", ctr.Status)

	return app.Listen(host)
}