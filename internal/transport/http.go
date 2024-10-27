package transport

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vndg-rdmt/paymentspod/internal/controller"
)

func NewHttp(ctr *controller.Fiber, host string) error {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(_ string) bool {
			return true
		},
		AllowHeaders:     "",
		AllowCredentials: true,
	}))
	app.Post("/api/payment", ctr.Payment)
	app.Get("/api/status/:id", ctr.Status)

	fmt.Println(host)

	return app.Listen(host)
}
