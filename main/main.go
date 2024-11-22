package main

import (
	"fmt"
	"log"
)

// Курси валют
var rates = map[string]float64{
	"USD": 1.0,           // 1 долар = 1 долар
	"UAH": 41.48,         // 1 долар = 41.48 гривень
	"EUR": 43.43 / 41.48, // 1 євро = 43.43 гривень
}

func main() {
	fmt.Println("Конвертер валют (USD, EUR, UAH)")
	fmt.Println("1. Показати доступні валюти")
	fmt.Println("2. Конвертувати валюту")
	fmt.Print("Виберіть опцію (1 або 2): ")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		showCurrencies()
	case 2:
		convertCurrency()
	default:
		fmt.Println("Неправильний вибір.")
	}
}

// Вивід доступних валют
func showCurrencies() {
	fmt.Println("Доступні валюти: USD, EUR, UAH")
}

// Конвертація валют
func convertCurrency() {
	var from, to string
	var amount float64

	fmt.Print("Введіть валюту для конвертації (USD, EUR, UAH): ")
	fmt.Scanln(&from)
	if _, ok := rates[from]; !ok {
		log.Fatalf("Валюта %s не підтримується.\n", from)
	}

	fmt.Print("Введіть суму: ")
	fmt.Scanln(&amount)

	fmt.Print("Введіть валюту, в яку конвертувати (USD, EUR, UAH): ")
	fmt.Scanln(&to)
	if _, ok := rates[to]; !ok {
		log.Fatalf("Валюта %s не підтримується.\n", to)
	}

	// Конвертація суми
	result := amount * rates[to] / rates[from]
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}
