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
		fmt.Println("Enter Integer number or X to exit : ")
		 line := scanner.Text()
		 if line == "X" {
			break // to stop writing enter X and press enter
		 }
		 num, _ := strconv.Atoi(line)
		 sli = append(sli, num)
		 // sort the slice
		 sort.Ints(sli)
	}
	fmt.Print(sli)




}
	
  
	