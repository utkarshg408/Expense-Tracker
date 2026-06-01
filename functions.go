package main

import (
	"fmt"
)

func addexpense() {
	var title, category, date string
	var amount float64

	fmt.Print("Enter the title of the expense : ")
	fmt.Scan(&title)
	fmt.Print("Enter the amount of the expense : ")
	fmt.Scan(&amount)
	fmt.Print("Enter the category of the expense : ")
	fmt.Scan(&category)
	fmt.Print("Enter the date of the expense : ")
	fmt.Scan(&date)

	var newexpense expense = expense{
		Title:    title,
		Amount:   amount,
		Category: category,
		Date:     date,
	}
	expenses = append(expenses, newexpense)
	fmt.Println("Expense added successfully !")

}

func viewallexpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses to show !")
		return
	}
	fmt.Println("All Expenses : ")
	for i := 0; i < len(expenses); i++ {
		fmt.Println("Title : ", expenses[i].Title)
		fmt.Println("Category : ", expenses[i].Category)
		fmt.Println("Amount : ", expenses[i].Amount)
		fmt.Println("Date : ", expenses[i].Date)
		fmt.Println()

	}
}

func viewtotalexpenses() {
	var total float64 = 0
	for i := 0; i < len(expenses); i++ {
		total += expenses[i].Amount
	}
	fmt.Printf("Total expenses: %.2f\n", total)
}
