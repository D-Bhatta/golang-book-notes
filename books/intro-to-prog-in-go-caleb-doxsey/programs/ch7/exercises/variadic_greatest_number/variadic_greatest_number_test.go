// variadic_greatest_number_test.go

package variadic_greatest_number

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

type Test struct {
	Id          uint64  `json:"id"`
	LargestNum  int64   `json:"largest"`
	Numbers     []int64 `json:"numbers"`
	ErrExpected error   `json:"err_expected"`
}

type Tests struct {
	Tests []Test `json:"tests"`
}

func TestGreatestNumber(t *testing.T) {

	testFileName := "test_cases.json"

	jsonFile, jsonFileErr := os.ReadFile(testFileName)

	if jsonFileErr != nil {
		log.Fatalf("Error reading JSON data from test data file: %s. \nError: %v", testFileName, jsonFileErr)
	}

	var tests Tests

	jsonUnmarshallErr := json.Unmarshal(jsonFile, &tests)

	if jsonUnmarshallErr != nil {
		log.Fatalf("Error unmarshalling JSON data from byte array. Error: %v", jsonUnmarshallErr)
	}

	for _, tc := range tests.Tests {
		t.Run(fmt.Sprintf("N%d", tc.Id), func(t *testing.T) {
			largest, err := GreatestNumber(tc.Numbers...)

			if err != tc.ErrExpected {
				t.Fatalf("Error in test ID %d: %v.\nExpected error is %v", tc.Id, err, tc.ErrExpected)
			}

			if largest != tc.LargestNum {
				t.Errorf("Given number sequence = %d.\nReturned largest number = %d.\nExpected largest number = %d.",
					tc.Numbers, largest, tc.LargestNum)
			}
		})
	}
}
