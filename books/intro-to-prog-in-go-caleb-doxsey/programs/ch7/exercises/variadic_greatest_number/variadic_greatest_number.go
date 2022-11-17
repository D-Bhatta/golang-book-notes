// variadic_greatest_number.go

package variadic_greatest_number

// find the greatest number in the list
func GreatestNumber(nums ...int64) (int64, error) {
    greatest := nums[0]
    for _, num := range nums {
        if num > greatest{
            greatest = num
        }
    }
	return greatest, nil
}
