package main

import (
	"encoding/csv"
	// "encoding/json"
	"github.com/j-larson/unmarshal/json"
	"fmt"
	"os"
	"time"
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
	NUM_SOURCE_RECORDS := len(csvRecords)
	fmt.Printf("Read %d records.\n", NUM_SOURCE_RECORDS)

	bytesArrs := make([][]byte, NUM_SOURCE_RECORDS)
	for i, r := range csvRecords {
		bytesArrs[i] = []byte(r[4])
	}

	NUM_TARGET_RECORDS := 1000
	targetArr := make([]interface{}, NUM_TARGET_RECORDS)
	fmt.Printf("Unmarshalling into %d records.\n", NUM_TARGET_RECORDS)

	NUM_ITERATIONS := 1000 * 1000;
	fmt.Printf("Unmarshalling %d times.\n", NUM_ITERATIONS)
	
	start := time.Now()
	for i := 0; i < NUM_ITERATIONS; i++ {
		srcIx := i % NUM_SOURCE_RECORDS
		dstIx := i % NUM_TARGET_RECORDS
		json.RecycleJson(targetArr[dstIx])
		err := json.Unmarshal(bytesArrs[srcIx], &targetArr[dstIx])
		if err != nil {
			fmt.Printf("Unable to unmarshal value %d: %v", i, err)
			return
		}
	}
	testDuration := time.Since(start)

	fmt.Printf("Test duration: %f s\n", testDuration.Seconds())
	//json.Report()
}

