package main

import (
	"fmt"
)

func main() {
	var uahAmount float64
	uahToUsdRate := 0.0024 // курс гривни к доллару (примерный)
	uahToEurRate := 0.0022 // курс гривни к евро (примерный)

	fmt.Print("Введите сумму в гривнах: ")
	fmt.Scan(&uahAmount)

	usdAmount := uahAmount * uahToUsdRate
	eurAmount := uahAmount * uahToEurRate

	fmt.Printf("Сумма в долларах: %.2f USD\n", usdAmount)
	fmt.Printf("Сумма в евро: %.2f EUR\n", eurAmount)
}
