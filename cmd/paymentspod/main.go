package main

import (
	"os"

	"github.com/vndg-rdmt/paymentspod/internal/controller"
	"github.com/vndg-rdmt/paymentspod/internal/service"
	"github.com/vndg-rdmt/paymentspod/internal/transport"
)

func main() {

	err := transport.NewHttp(
		controller.NewFiber(
			service.New(
				os.Getenv("TOCHKA_CUSTOMERCODE"),
				os.Getenv("TOCHKA_TOKEN"),
				os.Getenv("TOCHKA_URL"),
			),
		),
		os.Getenv("LISTEN_ADDR"),
	)
	if err != nil {
		panic(err)
	}
}
