package prime

import "fmt"

var noPrimeMap = make(map[int]bool)

// CountPrimeThree 素数统计方法3
func CountPrimeThree(n int) {
	var count int
	for i := 2; i < n; i++ {
		if _, ok := noPrimeMap[i]; !ok {
			for j := i * 2; j < n; j += i {
				noPrimeMap[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if _, ok := noPrimeMap[i]; !ok {
			count++
		}
	}
	fmt.Println("prime three = ", count)
}
