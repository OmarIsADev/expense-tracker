package cmd

import (
	"errors"

	"github.com/omarisadev/expense-tracker/internal"
	"github.com/omarisadev/expense-tracker/model"
	"github.com/spf13/cobra"
)

var description *string
var amount *float64

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an expense",
		Long: `add -d [description] -a [amount]
		
description: expense description
amount: expense amount`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return addCmd()
		},
	}

	description = cmd.Flags().StringP("description", "d", "", "Expense description")
	amount = cmd.Flags().Float64P("amount", "a", 0, "Expense amount")

	return cmd
}

func addCmd() error {
	if *description == "" || *amount == 0 {
		return errors.New("description and amount are required")
	}

	return internal.AddExpense(model.Expense{Description: *description, Amount: *amount})
}
