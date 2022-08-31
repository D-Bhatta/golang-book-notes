// divide_by_half.go
package divide_by_half

// returns the passed number divided in half and if it was even
func DivideByHalf(num int64) (int64, bool, error) {
    var half int64 = num/2
    var even bool = false
    if num % 2 == 0 {
        even = true
    } else {
        even = false
    }

    return half, even, nil
}