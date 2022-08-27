/*
Write a Bubble Sort program in Go. The program should prompt the user
to type in a sequence of up to 10 integers. The program should print
the integers out on one line, in sorted order, from least to greatest.

As part of this program, you should write a function called BubbleSort()
which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements
are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation
which swaps the position of two adjacent elements in the slice. You
should write a Swap() function which performs this operation. Your
Swap() function should take two arguments, a slice of integers and
an index value i which indicates a position in the slice. The Swap()
function should return nothing, but it should swap the contents of
the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	nums = make([]int, 0, 3)
	
)

func swap(num1 int, num2 int) {
	num1, num2 = num2, num1
}

func main() {
	// accept numbers from user
	fmt.Println("Enter Integer number or X to exit : ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		
		 line := scanner.Text()
		 if line == "X" {
			break // to stop writing enter X and press enter
		 }
		 num, _ := strconv.Atoi(line)
		 nums = append(nums, num)
	
		 fmt.Println("Enter Integer number or X to exit : ")
	}
	// sort numbers
	swapped := true
    for swapped {
        swapped = false
        for i := 1; i < len(nums); i++ {
            if nums[i-1] > nums[i] {
                nums[i], nums[i-1] = nums[i-1], nums[i]
                swapped = true
            }
        }
    }
	fmt.Println(nums)
}