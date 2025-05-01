# Expense Tracker
https://roadmap.sh/projects/expense-tracker
A simple command line application to track expenses.

## Installation

Add to path:
```bash
go install
expense-tracker [command]
```

## Build & run

Linux:
```bash
go build .
./tasks-cli
```
Windows: 
```bash
go build .
./tasks-cli.exe 
```

## Usage

### Add Expense

`expense-tracker add -d <description> -a <amount>`

Example: `expense-tracker add -d "Lunch" -a 10.99`

### List Expenses

`expense-tracker list`

### Delete Expense

`expense-tracker delete -i <id>`

Example: `expense-tracker delete -i 1`

### Summarize Expenses

`expense-tracker summary -m <month>`

Example: `expense-tracker summary -m jan`
