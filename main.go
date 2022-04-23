package main

import "fmt"

func main() {
	var running bool = true
	fmt.Println("Welcome To SuperCalc")
	for running {
		var operator string
		fmt.Scan(&operator)
		if operator == "add" {
			var value1 int
			var value2 int
			fmt.Println("Enter First Value: ")
			fmt.Scan(&value1)
			fmt.Println("Enter Second Value: ")
			fmt.Scan(&value2)
			var result int = add(value1, value2)
			fmt.Printf("%v + %v = %v\n", value1, value2, result)
		} else if operator == "sub" {
			var value1 int
			var value2 int
			fmt.Println("Enter First Value: ")
			fmt.Scan(&value1)
			fmt.Println("Enter Second Value: ")
			fmt.Scan(&value2)
			var result int = sub(value1, value2)
			fmt.Printf("%v - %v = %v\n", value1, value2, result)
		} else if operator == "mult" {
			var value1 int
			var value2 int
			fmt.Println("Enter First Value: ")
			fmt.Scan(&value1)
			fmt.Println("Enter Second Value: ")
			fmt.Scan(&value2)
			var result int = mult(value1, value2)
			fmt.Printf("%v * %v = %v\n", value1, value2, result)
		} else if operator == "div" {
			var value1 int
			var value2 int
			fmt.Println("Enter First Value: ")
			fmt.Scan(&value1)
			fmt.Println("Enter Second Value: ")
			fmt.Scan(&value2)
			var result int = div(value1, value2)
			fmt.Printf("%v / %v = %v\n", value1, value2, result)
		} else if operator == "quit" || operator == "exit" {
			fmt.Println("Thanks for using SuperCalc!")
			break
		} else {
			fmt.Println("Argument unknown, please try again")
		}
	}
}

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mult(x int, y int) int {
	return x * y
}

func div(x int, y int) int {
	return x / y
}
