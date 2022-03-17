package main

import "fmt"

func main() {
	var n, a, r int

	fmt.Scanln(&n, &a, &r)
	i := 1
	fmt.Print(0)
	for i <= n {
		fmt.Print(" + ", a*i*r)
		i++
	}
}
