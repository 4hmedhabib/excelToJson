package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

// ConvertSheetToJSON reads an Excel sheet and converts it to JSON
func ConvertSheetToJSON(excelFile, sheetName, outputFile string) error {
	// Open the Excel file
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// Check if the sheet exists
	sheets := f.GetSheetList()
	found := false
	for _, s := range sheets {
		if s == sheetName {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("sheet %s not found in file %s", sheetName, excelFile)
	}

	// Read all rows from the sheet
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("failed to read sheet %s: %v", sheetName, err)
	}

	if len(rows) < 2 {
		return fmt.Errorf("not enough data in sheet %s", sheetName)
	}

	// First row is used as JSON keys (headers)
	headers := rows[0]
	var jsonData []map[string]string

	// Convert each row into a key-value pair
	for _, row := range rows[1:] {
		rowMap := make(map[string]string)
		for i, cell := range row {
			if i < len(headers) {
				rowMap[headers[i]] = cell
			}
		}
		jsonData = append(jsonData, rowMap)
	}

	// Convert to JSON format
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %v", err)
	}

	// Save JSON file
	err = os.WriteFile(outputFile, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	fmt.Printf("Sheet '%s' converted to %s\n", sheetName, outputFile)
	return nil
}

func main() {
	// Define command-line flags
	excelPath := flag.String("path", "", "Path to the Excel file")
	sheetName := flag.String("sheet", "", "Sheet name to convert")
	outputFile := flag.String("output", "", "Output JSON file name")

	flag.Parse()

	// Validate input arguments
	if *excelPath == "" || *sheetName == "" || *outputFile == "" {
		log.Fatal("Usage: go run main.go -path <excel_file.xlsx> -sheet <SheetName> -output <output.json>")
	}

	// Convert sheet to JSON
	err := ConvertSheetToJSON(*excelPath, *sheetName, *outputFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
