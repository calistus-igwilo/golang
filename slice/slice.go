/*
Write a program which prompts the user to enter integers and stores
the integers in a sorted slice. The program should be written as a
loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop,
the program prompts the user to enter an integer to be added to the
slice. The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order. The slice must
grow in size to accommodate any number of integers which the user
decides to enter. The program should only quit (exiting the loop)
when the user enters the character 'X' instead of an integer.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var num int

var sli = make([]int, 0, 3) // create a slice with lenght 3, start from zero
func main() {
	fmt.Println("Enter Integer number or X to exit : ")
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		
		 line := scanner.Text()
		 if line == "X" {
			break // to stop writing enter X and press enter
		 }
		 num, _ := strconv.Atoi(line)
		 sli = append(sli, num)
		 // sort the slice
		 sort.Ints(sli)
		 fmt.Print(sli)
		 fmt.Println("Enter Integer number or X to exit : ")
	}
	
}