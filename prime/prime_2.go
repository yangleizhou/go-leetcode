package prime

import "fmt"

// CountPrimeTwo 素数统计方法2
func CountPrimeTwo(n int) {
	var count int
	for i := 2; i < n; i++ {
		if isPrimeTwo(i) {
			count++
		}
	}
	fmt.Println("prime two = ", count)
}

func isPrimeTwo(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
