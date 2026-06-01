package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Expense Tracker !")
	var choice int
	for choice != 4 {
		fmt.Println("Select the option you want to perform : ")
		fmt.Println("1.add an expense ")
		fmt.Println("2.view all expenses ")
		fmt.Println("3.view total expenses ")
		fmt.Println("4.exit ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			addexpense()
		case 2:
			viewallexpenses()
		case 3:
			viewtotalexpenses()
		case 4:
			fmt.Println("Exiting the application !")
		default:
			fmt.Println("Invalid choice ! Please select a valid option.")

		}
	}

}
