package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Введення виразу
	fmt.Print("Введіть вираз (наприклад, 3.5 + 2 * 6 / 3): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Розділення виразу на числа і оператори
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		log.Fatal("Помилка: порожній вираз")
	}

	var result float64
	var currentOperator string

	for i, token := range tokens {
		// Парсинг чисел
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			if i == 0 {
				// Якщо перше число, встановлюємо його як результат
				result = num
			} else {
				// Застосування оператора
				switch currentOperator {
				case "+":
					result += num
				case "-":
					result -= num
				case "*":
					result *= num
				case "/":
					if num == 0 {
						log.Fatal("Помилка: ділення на нуль")
					}
					result /= num
				default:
					log.Fatal("Помилка: невідомий оператор")
				}
			}
		} else {
			// Якщо це не число, тоді перевіряємо, чи це оператор
			if token == "+" || token == "-" || token == "*" || token == "/" {
				currentOperator = token
			} else {
				log.Fatal("Помилка: введено невірний символ")
			}
		}
	}

	fmt.Printf("Результат: %.2f\n", result)
}
