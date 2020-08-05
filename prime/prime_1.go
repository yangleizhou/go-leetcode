package prime

import "fmt"

// CountPrimeOne 素数统计方法1
func CountPrimeOne(n int) {
	var count int
	for i := 2; i < n; i++ {
		if isPrimeOne(i) {
			count++
		}
	}
	fmt.Println("prime one = ", count)
}

func isPrimeOne(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
