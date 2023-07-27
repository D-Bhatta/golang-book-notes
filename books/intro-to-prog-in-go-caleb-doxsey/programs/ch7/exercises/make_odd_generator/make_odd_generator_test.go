// make_odd_generator_test.go

package make_odd_generator

import (
	"fmt"
	"testing"
)

func TestGenerateOddNumList(t *testing.T) {

	odds := []struct {
		ID          uint64
		StartNum    uint64
		ListSize    uint64
		NumList     []uint64
		ErrExpected error
	}{
		{1, 1, 11, []uint64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, nil},
        {2, 3, 6, []uint64{3, 5, 7, 9, 11, 13}, nil},
        {3, 11, 5, []uint64{11, 13, 15, 17, 19}, nil},
        {4, 13, 0, []uint64{}, nil},
	}

	for _, tc := range odds {
		t.Run(fmt.Sprintf("N%d", tc.ID), func(t *testing.T) {
			num_list, err := GenerateOddNumList(tc.StartNum, tc.ListSize)
			if err != tc.ErrExpected {
				t.Fatalf("Error in test %v.\nExpected error is %v.", err, tc.ErrExpected)
			}
			if len(num_list) != len(tc.NumList) {
				t.Fatalf("Returned length of number list = %d, expected length of number list = %d for test case ID = %d",
					len(num_list), len(tc.NumList), tc.ID)
			}
            for idx, _ := range tc.NumList {
                if num_list[idx] != uint(tc.NumList[idx]){
                    t.Errorf("Returned list %v doesn't match expected list %v at idx = %d", num_list, tc.NumList, idx)
                }
            }
		})
	}

}
