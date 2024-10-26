package main

import (
	"fmt"
)

func main() {
	var uahAmount float64
	uahToUsdRate := 0.0024 // курс гривні к доллару (приблизний)
	uahToEurRate := 0.0022 // курс гривні к евро (приблизний)

	fmt.Print("Введіть суму в гривнях: ")
	fmt.Scan(&uahAmount)

	usdAmount := uahAmount * uahToUsdRate
	eurAmount := uahAmount * uahToEurRate

	fmt.Printf("Сума в долларах: %.2f USD\n", usdAmount)
	fmt.Printf("Сума в евро: %.2f EUR\n", eurAmount)
}
