package main

import (
	"fmt"
	"strconv"
)

func readInput() ([]int, error) {
	var res []int
	for i := 0; i < 10; i++ {
		var s string
		fmt.Println("insert integer:")
		_, err := fmt.Scanln(&s)
		if err != nil {
			return nil, err
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			break
		}
		res = append(res, x)
	}
	return res, nil
}

func swap(a []int, i int) {
	tmp := a[i]
	a[i] = a[i+1]
	a[i+1] = tmp
}

func bubbleSort(a []int) {
	for {
		swapped := false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				swap(a, i)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	ints, err := readInput()
	if err != nil {
		panic(err)
	}
	bubbleSort(ints)

	fmt.Println(ints)
}
