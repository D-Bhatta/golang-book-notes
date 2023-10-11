// recursive_fibonacci.go

package recursivefibonacci

func RecursiveFibonacci(n uint64) (uint64) {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2))
}
