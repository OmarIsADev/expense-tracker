package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/omarisadev/expense-tracker/model"
)

func GetFilePath(fileName string) string {
	cwd, err := os.Getwd()
	if os.IsNotExist(err) {
		fmt.Println("Error Reading file path", err)
	}

	return path.Join(cwd, fileName)
}

func ReadExpensesFile() ([]model.Expense, error) {
	filePath := GetFilePath("expense.json")

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("Save File does not exist, Creating new one.")

		file, err := os.Create(filePath)
		os.WriteFile(file.Name(), []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Error creating file", err)
			return nil, err
		}

		defer file.Close()
		fmt.Println("Save file created successfully")

		return []model.Expense{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var expenses []model.Expense
	err = json.NewDecoder(file).Decode(&expenses)
	if err != nil {
		fmt.Println("Error reading save file", err)
		return nil, err
	}

	return expenses, nil
}

func WriteExpensesFile(expenses []model.Expense) error {
	filePath := GetFilePath("expense.json")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating save file", err)
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(expenses)
	if err != nil {
		fmt.Println("Error writing save file", err)
		return err
	}

	return nil
}
