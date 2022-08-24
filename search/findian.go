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
	"bufio"
	"fmt"
	"os"
	"strings"
)

var str string 

func main() {
	fmt.Println("Enter string, please : ")
	reader := bufio.NewReader(os.Stdin)

	// ReadString will block until the delimiter is entered
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading string. Please try again", err)
		return
	}
	
	// remove the delimeter from the string
	str = strings.TrimSuffix(str, "\n")
	
	// convert string to lowercase 
	strLower := strings.ToLower(str)

	if strings.HasPrefix(strLower, "i") && strings.HasSuffix(strLower, "n") && strings.Contains(strLower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}