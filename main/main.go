package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var input1, input2, operator string

	//введення першого числа
	fmt.Print("введіть перше число: ")
	fmt.Scanln(&input1)

	// перевірка першого числа на ціле
	a, err := strconv.Atoi(input1)
	if err != nil {
		log.Fatal("Помидка: введено не ціле число")
	}

	// введення оператора
	fmt.Print("введіть оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	//перевірка оператора
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		log.Fatal("помилка: введено не оператор")
	}
	//  перевірка другого числа на ціле
	b, err := strconv.Atoi(input2)
	if err != nil {
		log.Fatal("помилка: введено не ціле число")
	}

	// виконання операції
	switch operator {
	case "+":
		fmt.Printf("результат: %d\n", a+b)
	case "-":
		fmt.Printf("результат: %d\n", a-b)
	case "*":
		fmt.Printf("результат: %d\n", a*b)
	case "/":
		if b == 0 {
			log.Fatal("помилка: ділення на нуль")
		}
		fmt.Printf("результат: %d\n", a/b)
	default:
		log.Fatal("помилка: невідома операція")
	}

}
