package cmd

import (
	"github.com/omarisadev/expense-tracker/internal"
	"github.com/spf13/cobra"
)

var month *string

func NewSummaryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "summary -m",
		Short: "Summary of expenses",
		Long: `summary -m [month]
		
month: all, jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return summaryCmd()
		},
	}

	month = cmd.Flags().StringP("month", "m", "all", "Month")

	return cmd
}

func summaryCmd() error {
	return internal.SummaryExpenses(*month)
}
