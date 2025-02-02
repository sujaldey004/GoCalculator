package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Calculator")
	fmt.Println("-------------")

	var infCalculator bool = true
	var choice string
	for infCalculator {
		var num1, num2 int
		fmt.Print("Enter first number : ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number : ")
		fmt.Scan(&num2)

		operations := []string{"1) Addition", "2) Subtraction", "3) Multiplication", "4) Division"}
		for i := range operations {
			fmt.Printf("%v\n", operations[i])
		}

		fmt.Println("Choose operation you want to perform")
		var userInput int
		var result float64
		fmt.Scan(&userInput)
		switch userInput {
		case 1:
			result = float64(num1 + num2)
			fmt.Printf("Result of Addition is : %v", result)
		case 2:
			result = float64(num1 - num2)
			fmt.Printf("Result of Subtraction is : %v", result)
		case 3:
			result = float64(num1 * num2)
			fmt.Printf("Result of Multiplication is : %v", result)
		case 4:
			result = float64(num1 / num2)
			fmt.Printf("Result of the Division is : %v", result)
		}
		fmt.Print("\nDo you want to continue 'y' for yes and 'n' for no : ")
		fmt.Scan(&choice)
		if choice == "y" {
			infCalculator = true
		} else if choice == "n" {
			infCalculator = false
		} else {
			fmt.Println("Invalid choice!")
			break
		}
	}
}
