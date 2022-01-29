package main

import "fmt"

func sn(x int, n int) int {
	return (x * n * (n + 1)) / 2
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCases > 0 {
		testCases = testCases - 1
		var n int
		fmt.Scanln(&n)
		n = n - 1
		s3 := sn(3, n/3)
		s5 := sn(5, n/5)
		s15 := sn(15, n/15)

		fmt.Println(s3 + s5 - s15)
	}
}
