package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("combined.csv")
	if err != nil {
		fmt.Printf("Unable to open CSV file: %v\n", err)
		return
	}

	csvReader := csv.NewReader(file)
	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to read CSV file: %v\n", err)
		return
	}


	fmt.Printf("Read %d records.\n", len(csvRecords))

	for i, r := range csvRecords {
		jsonString := r[4]

		bytes := []byte(jsonString)
		var target interface{}

		err := json.Unmarshal(bytes, &target)
		if err != nil {
			fmt.Printf("Unable to unmarshal value %d: %s: %v", i, jsonString, err)
			return
		}
	}

	fmt.Printf("Done.\n")
}

