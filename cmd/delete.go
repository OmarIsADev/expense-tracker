package cmd

import (
	"github.com/omarisadev/expense-tracker/internal"
	"github.com/spf13/cobra"
)

var id *int

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete --id [id]",
		Short: "Delete an expense",
		Long: `delete --id [id]
		
id: expense ID`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteCmd()
		},
	}

	id = cmd.Flags().IntP("id", "i", 0, "Expense ID")
	return cmd
}

func deleteCmd() error {
	return internal.DeleteExpense(uint16(*id))
}
