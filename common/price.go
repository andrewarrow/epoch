package common

import (
	"fmt"
	"strconv"
)

func SimplePrice(s any) string {
	amount := fmt.Sprintf("%v", s)
	price, _ := strconv.Atoi(amount)
	if price == 0 {
		return "0.00"
	}
	if len(amount) < 3 {
		return fmt.Sprintf("00.%s", amount)
	}
	dollars := amount[0 : len(amount)-2]
	//dollarsInt, _ := strconv.ParseInt(dollars, 10, 64)
	//dollars = util.IntComma(dollarsInt)
	cents := amount[len(amount)-2:]
	return fmt.Sprintf("%s.%s", dollars, cents)
}
