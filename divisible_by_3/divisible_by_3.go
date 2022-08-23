package main

import "fmt"

func main() {
	// print all the numbers from 1 to 100 that are
	// divisible by 3

	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
		
	}
}