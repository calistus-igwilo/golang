/*
Write a program which prompts the user to enter a string. The program searches through
the entered string for the characters 'i', 'a', and 'n'. The program should print 'Found!'
if the entered string starts with the character 'i', ends with the character 'n', and
contains the character 'a'. The program should print 'Not Found!' otherwise. The program
should not be case-sensitive, so it does not matter if the characters are upper-case or
lower-case
*/

package main

import (
	"fmt"
	"strings"
)

var str string 

func main() {
	fmt.Println("Enter text, please : ")
	fmt.Scanf("%s",  &str)
	strLower := strings.ToLower(str)

	if strLower[0] == 'i' && strLower[len(strLower)-1] == 'n' && strings.Contains(strLower, "a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}