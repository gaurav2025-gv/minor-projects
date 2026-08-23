package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Expense struct {
	ID          int
	Amount      float64
	Category    string
	Description string
}

var expenses []Expense
var nextID = 1

func addExpense(reader *bufio.Reader) {
	fmt.Print("Amount: ")
	amountInput, _ := reader.ReadString('\n')
	amountInput = strings.TrimSpace(amountInput)

	amount, err := strconv.ParseFloat(amountInput, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount!")
		return
	}

	fmt.Print("Category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	if category == "" {
		fmt.Println("Category cannot be empty!")
		return
	}

	fmt.Print("Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	expense := Expense{
		ID:          nextID,
		Amount:      amount,
		Category:    category,
		Description: description,
	}

	expenses = append(expenses, expense)
	nextID++

	fmt.Println("Expense added!")
}

func listExpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Println("\n===== EXPENSES =====")

	for _, expense := range expenses {
		fmt.Printf(
			"%d. ₹%.2f | %s | %s\n",
			expense.ID,
			expense.Amount,
			expense.Category,
			expense.Description,
		)
	}
}

func totalExpenses() {
	var total float64

	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Printf("Total Expenses: ₹%.2f\n", total)
}

func categoryExpenses(reader *bufio.Reader) {
	fmt.Print("Enter category: ")

	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	var total float64

	for _, expense := range expenses {
		if strings.EqualFold(expense.Category, category) {
			total += expense.Amount
		}
	}

	fmt.Printf(
		"%s Expenses: ₹%.2f\n",
		category,
		total,
	)
}

func deleteExpense(reader *bufio.Reader) {
	fmt.Print("Enter Expense ID: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	for i := range expenses {
		if expenses[i].ID == id {

			expenses = append(
				expenses[:i],
				expenses[i+1:]...,
			)

			fmt.Println("Expense deleted!")
			return
		}
	}

	fmt.Println("Expense not found.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== EXPENSE TRACKER =====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. List Expenses")
		fmt.Println("3. Total Expenses")
		fmt.Println("4. Show Category Expenses")
		fmt.Println("5. Delete Expense")
		fmt.Println("6. Exit")

		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid choice!")
			continue
		}

		switch choice {

		case 1:
			addExpense(reader)

		case 2:
			listExpenses()

		case 3:
			totalExpenses()

		case 4:
			categoryExpenses(reader)

		case 5:
			deleteExpense(reader)

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}