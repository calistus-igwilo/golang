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
	"strings"
)

var (
	//nums = make([]int, 0, 3)
	
)
/* func ReadInput(){
	fmt.Println("Enter the numbers, seperate each by a space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := scanner.Text()
} */

func StringToInteger(str string) []int {
    var nums []int
    for _, char := range strings.Fields(str) {
        num, err := strconv.Atoi(char)
		if err != nil {
			return nil
		}
        nums = append(nums, num)
    }
    return nums
}


func Swap(nums []int, i int) {
	// Swap values at index i and index i+1
	nums[i], nums[i+1] = nums[i+1], nums[i]
}


func BubbleSort(arr []int){
	swapped := true
    for swapped {
        swapped = false
        for i := 0; i < len(arr)-1; i++ {
            if arr[i] > arr[i+1] {
                Swap(arr, i)
                swapped = true
            }
        }
    }
}


func main() {
	// accept numbers from user
	fmt.Println("Enter the numbers, seperate each by a space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	// convert input string to integer array
	nums := StringToInteger(scanner.Text()) 

	// sort the numbers
	BubbleSort(nums)

	// print sorted numbers on one line
	fmt.Println(nums)
}