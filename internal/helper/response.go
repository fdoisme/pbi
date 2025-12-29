package helper

import (
	"github.com/gofiber/fiber/v2"
)

// TODO : make helper response
type BaseResponse[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    T      `json:"data"`
}
type DataWithPagination[T any] struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Data  T   `json:"data"`
}

func GoodResponse[T any](ctx *fiber.Ctx, message string, data T) {
	response := BaseResponse[T]{
		Status:  false,
		Message: message,
		Data:    data,
	}
	ctx.Status(fiber.StatusOK).JSON(response)
}

func BadResponse[T any](ctx *fiber.Ctx, statusCode int, message string, err any) {
	response := BaseResponse[T]{
		Status:  false,
		Message: message,
		Errors:  err,
	}
	ctx.Status(statusCode).JSON(response)
}
