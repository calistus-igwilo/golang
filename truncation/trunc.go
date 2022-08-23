/*
A program which prompts the user to enter a floating point number and prints the
integer which is a truncated version of the floating point number that was entered
*/

package main

import "fmt"

var floatNum float64
func main(){
	fmt.Println("Enter a float value : ")
	fmt.Scanf("%f", &floatNum)
	fmt.Printf("Truncated float: %.0f ", floatNum)

}