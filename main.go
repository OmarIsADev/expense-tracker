package main

import (
	"github.com/omarisadev/expense-tracker/cmd"
)

func main() {
	// const padding = 3
	// w := tabwriter.NewWriter(os.Stdout, 1, 1, padding, ' ', 0)
	// fmt.Fprintf(w, "ID\tDescription\tAmount\n")
	// fmt.Fprintf(w, "1\tTest\t5\n")
	// fmt.Fprintf(w, "2\tSSUkaa\t10\n")
	// w.Flush()
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
