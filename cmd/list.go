package cmd

import (
	"github.com/omarisadev/expense-tracker/internal"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List expenses",
		Long:  `list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listCmd()
		},
	}

	return cmd
}

func listCmd() error {
	return internal.ListExpenses()
}
