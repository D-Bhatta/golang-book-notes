// divide_by_half_test.go

package divide_by_half

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

type Test struct {
    Id          uint64 `json:"test_id"`
	Num         int64 `json:"num"`
	HalfNum     int64 `json:"half_num"`
	Even        bool  `json:"even"`
	ErrExpected error `json:"err_expected"`
}

type Tests struct {
	Tests []Test `json:"tests"`
}

func TestDivideByHalf(t *testing.T) {

	testFileName := "test_cases.json"

	jsonFile, jsonFileErr := os.ReadFile(testFileName)
	// because it should be used instead of ioutil

	if jsonFileErr != nil {
		log.Fatalf("Error reading data from test file: %s.\nError: %v", testFileName, jsonFileErr)
	}

	var tests Tests

	jsonUnmarshallErr := json.Unmarshal(jsonFile, &tests)

	if jsonUnmarshallErr != nil {
		log.Fatalf("Error unmarshalling data from bytearray. Error: %v", jsonUnmarshallErr)
	}

	for _, tc := range tests.Tests {
		t.Run(fmt.Sprintf("N%d", tc.Id), func(t *testing.T) {
			halfNum, even, err := DivideByHalf(tc.Num)
			if err != tc.ErrExpected {
				t.Fatalf("Error in test %v. Expected error is %v", err, tc.ErrExpected)
			}
			if halfNum != tc.HalfNum || even != tc.Even {
				t.Errorf("Given tc.Num = %d, Returned halfNum = %d, even = %t. Expected halfNum = %d, even = %t\n",
				tc.Num, halfNum, even, tc.HalfNum, tc.Even)
			}
		})
	}
}
