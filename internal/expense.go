package internal

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/omarisadev/expense-tracker/model"
)

func ListExpenses() error {
	expenses, err := ReadExpensesFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		return fmt.Errorf("no expenses found, try to add using add command")
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tDescription\tAmount\tDate")
	for _, expense := range expenses {
		fmt.Fprintf(w, "%d\t%s\t$%.2f\t%s\n", expense.ID, expense.Description, expense.Amount, expense.CreatedAt.Format("2006-01-02"))
	}
	w.Flush()

	return nil
}

func AddExpense(expense model.Expense) error {
	expenses, err := ReadExpensesFile()
	if err != nil {
		return err
	}

	expense.ID = uint16(len(expenses) + 1)
	expense.CreatedAt = time.Now()
	expenses = append(expenses, expense)

	err = WriteExpensesFile(expenses)
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully ID(%d)\n", expense.ID)

	return nil
}

func DeleteExpense(id uint16) error {
	expenses, err := ReadExpensesFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		return fmt.Errorf("no expenses found")
	}

	if len(expenses) < int(id) {
		return fmt.Errorf("expense with ID(%d) does not exist", id)
	}

	var newExpenses []model.Expense

	for i, expense := range expenses {
		if expense.ID != id {
			if expense.ID > id {
				expense.ID -= 1
			} else {
				expense.ID = uint16(i + 1)
			}
			newExpenses = append(newExpenses, expense)
		}
	}

	err = WriteExpensesFile(newExpenses)
	if err != nil {
		return err
	}

	fmt.Printf("Expense deleted successfully ID(%d)\n", id)

	return nil
}

func SummaryExpenses(month string) error {
	expenses, err := ReadExpensesFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		return fmt.Errorf("no expenses found, try to add using add command")
	}

	var total float64
	for _, expense := range expenses {
		if month == "all" || strings.EqualFold(expense.CreatedAt.Month().String()[0:3], month[0:3]) {
			total += expense.Amount
		}
	}

	if total == 0 {
		return fmt.Errorf("no expenses found for month(%s)", month)
	}

	fmt.Printf("Total Expenses: $%.2f\n", total)
	return nil
}
