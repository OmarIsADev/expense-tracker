package main

import (
	"github.com/omarisadev/expense-tracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
