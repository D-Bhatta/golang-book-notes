// make_odd_generator.go

package make_odd_generator

// generate odd number
func GenerateOdd(start uint) func() uint {
    num := uint(start)
	return func () (ret uint) {
        ret = num
        num += 2
        return
    }
}

func GenerateOddNumList(start uint64, size uint64) ([]uint, error) {
    if size == 0 {
        return []uint{}, nil
    }
    nextNum := GenerateOdd(uint(start))
    var numList []uint
    for i:=uint64(0); i<size; i++ {
        numList = append(numList, nextNum())
    }
    return numList, nil
}