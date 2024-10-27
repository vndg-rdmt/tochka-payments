package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vndg-rdmt/paymentspod/internal/service"
)

type MakePayment struct {
	Amount  uint64 `json:"amount"`
	Purpose string `json:"string"`
}

type MakePaymentResponse struct {
	Id string `json:"id"`
}

func NewFiber(srv service.Service) *Fiber {
	return &Fiber{
		service: srv,
	}
}

type Fiber struct {
	service service.Service
}

func (f *Fiber) Payment(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), time.Minute)
	defer cancel()

	var payload MakePayment

	dec := json.NewDecoder(bytes.NewBuffer(c.Body()))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	host := string(c.Request().Host())
	paymentId, err := f.service.Payment(ctx, host, host, payload.Amount, payload.Purpose)
	if err != nil {
		if tochkaerr, ok := err.(*service.TochkaError); ok {
			if code, converted := strconv.Atoi(tochkaerr.Code); converted == nil {
				return c.Status(code).JSON(tochkaerr)
			}
			return c.Status(fiber.StatusBadRequest).JSON(tochkaerr)
		}
	}

	return c.Status(fiber.StatusOK).JSON(MakePaymentResponse{
		Id: paymentId,
	})
}

func (f *Fiber) Status(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), time.Minute)
	defer cancel()

	resp, err := f.service.Status(ctx, c.Params("id", ""))
	if err != nil {
		return c.SendStatus(fiber.StatusServiceUnavailable)
	}

	c.Set("content-type", "application/json")
	return c.Status(fiber.StatusOK).Send(resp)
}
