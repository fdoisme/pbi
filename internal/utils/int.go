package utils

import (
	"fmt"
	"strconv"
)

func StrToUint(data string) uint64 {
	result, errConv := strconv.ParseUint(data, 10, 32)
	if errConv != nil {
		fmt.Println("Error StrToUint:", errConv)
	}
	return result
}
func StrToInt(data string) int {
	result, errConv := strconv.Atoi(data)
	if errConv != nil {
		fmt.Println("Error StrToUint:", errConv)
	}
	return result
}
