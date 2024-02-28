package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func kalkulator(name string) {

	menu := true
	var input string
	operator := []string{"+", "-", "*", "/"}
	firstNumber := 0
	secondNumber := 0
	result := 0

	for menu {
		clearScreen()
		fmt.Printf("Hello %s, Welcome To Calculator App\n", name)
		fmt.Println("+, -, *, /")
		fmt.Print("Enter Operator: ")
		fmt.Scan(&input)

		if !govalidator.IsIn(input, operator...) {
			fmt.Println("Invalid Operator!")
			sleepScreen(2)
			continue
		}

		clearScreen()

		fmt.Print("Enter First Number : ")
		fmt.Scan(&firstNumber)
		fmt.Print("Enter Second Number : ")
		fmt.Scan(&secondNumber)

		switch input {
		case "+":
			result = addition(firstNumber, secondNumber)
		case "-":
			result = substraction(firstNumber, secondNumber)
		case "/":
			result = division(firstNumber, secondNumber)
		case "*":
			result = multiplication(firstNumber, secondNumber)
		}

		fmt.Printf("Yeay Congrats %v, The result was %v", name,result)
		sleepScreen(5)
	}
}

func addition(first int, second int) int {
	return first + second
}

func substraction(first int, second int) int {
	return first - second
}

func division(first int, second int) int {
	return first / second
}

func multiplication(first int, second int) int {
	return first * second
}
