package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// FileNotFoundError represents a custom error for missing files.
type FileNotFoundError struct {
	Filename string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found: %s", e.Filename)
}

// CSVError represents a custom error for CSV parsing issues.
type CSVError struct {
	Message string
}

func (e *CSVError) Error() string {
	return fmt.Sprintf("CSV error: %s", e.Message)
}

// CalculationError represents a custom error for calculation issues.
type CalculationError struct {
	Row     int
	Message string
}

func (e *CalculationError) Error() string {
	return fmt.Sprintf("Calculation error in row %d: %s", e.Row, e.Message)
}

// CalculateRowSum calculates the sum of numbers in a CSV row.
func CalculateRowSum(row []string) (int, error) {
	sum := 0

	for i, cell := range row {
		num, err := strconv.Atoi(cell)
		if err != nil {
			return 0, &CalculationError{Row: i + 1, Message: "Invalid numeric data"}
		}
		sum += num
	}

	return sum, nil
}

// ProcessCSVFile reads and processes a CSV file.
func ProcessCSVFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, &FileNotFoundError{Filename: filename}
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return 0, &CSVError{Message: "Failed to parse CSV data"}
	}

	totalSum := 0

	for i, row := range rows {
		rowSum, err := CalculateRowSum(row)
		if err != nil {
			return 0, err
		}
		totalSum += rowSum
	}

	return totalSum, nil
}

func main() {
	filename := "data.csv" // Replace with the actual CSV file name

	totalSum, err := ProcessCSVFile(filename)
	if err != nil {
		switch e := err.(type) {
		case *FileNotFoundError:
			fmt.Println(e.Error())
		case *CSVError:
			fmt.Println(e.Error())
		case *CalculationError:
			fmt.Println(e.Error())
		default:
			fmt.Println("Unknown error:", err)
		}
		return
	}

	fmt.Println("Total sum of numbers:", totalSum)
}
