package utils

import (
	"fmt"
	"log"
	"time"
	"tugas_akhir_example/internal/helper"

	"github.com/gofiber/fiber/v2"
)

// TODO : make function parsing date
func ParseDateToGoTime(date string, err *helper.ErrorStruct) time.Time {
	tanggalLahir, errTime := time.Parse("02/01/2006", date)
	if errTime != nil {
		log.Println(errTime)
		fmt.Println("[PARSETIME]", errTime)
		err.Code = fiber.StatusBadRequest
		err.Err = errTime
		return time.Time{}
	}
	fmt.Println("[PARSETIME]", tanggalLahir)
	return tanggalLahir
}

func ParseDateToStr(date time.Time) string {
	return date.Format("2006-01-02")
}
