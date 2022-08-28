package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, index int) {
	tmp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = tmp
}

func main() {
	s := make([]int, 0, 10)
	cnt := 0
	for {
		var input string
		fmt.Println("enter a integer or 'X' to execute sorting a slice")
		fmt.Scanf("%s", &input)
		if "X" == strings.ToUpper(input) {
			BubbleSort(s)
			fmt.Println(s)
			break
		}
		if cnt > 9 {
			fmt.Println("can not enter a integer anymore")
		} else {
			num, _ := strconv.Atoi(input)
			s = append(s, num)
			cnt++
		}
	}
}
