package main

import "fmt"

var floatNum float64
func main(){
	fmt.Println("Enter a float value : ")
	fmt.Scanf("%f", &floatNum)
	fmt.Printf("%.0f", floatNum)

}