package prime

import "fmt"

var noPrimeMap4 = make(map[int]bool)

func CountPrimeFour(n int) {
	var count int
	for i := 2; i*i < n; i++ {
		if _, ok := noPrimeMap4[i]; !ok {
			for j := i * i; j < n; j += i {
				noPrimeMap4[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if _, ok := noPrimeMap4[i]; !ok {
			count++
		}
	}
	fmt.Println("prime four = ", count)
}
