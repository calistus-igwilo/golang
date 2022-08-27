/*
Write a program which reads information from a file and represents
it in a slice of structs. Assume that there is a text file which
contains a series of names. Each line of the text file has a first
name and a last name, in that order, separated by a single space on
the line.

Your program will define a name struct which has two fields, fname
for the first name, and lname for the last name. Each field will be
a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in
the file. Each struct created will be added to a slice, and after
all lines have been read from the file, your program will have a
slice containing one struct for each line in the file. After reading
all lines from the file, your program should iterate through your
slice of structs and print the first and last names found in each
struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct{
	fname string
	lname string 
}

func TruncateString(str string, length int) (truncated string) {
	// supports Japanese characters
    if length <= 0 {
        return
    }
    for i, char := range str {
        if i >= length {
            break
        }
        truncated += string(char)
    }
    return
}

//var fileName string
var textFile string 
var name1 name
var names = make([]name, 0, 3) // create a slice with lenght 3, start at 0 


func main() {
	const fileName = "test.txt"
	fmt.Println("Enter the name of text file: ")
	fmt.Scanf("%s", &textFile)

	file, err := os.Open(textFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		firstName, lastName := line[0], line[1]

		// Limit fristName and lastName to 20 characters
		firstName = TruncateString(firstName, 20)
		lastName = TruncateString(lastName, 20)

		// struct to hold each firstname and lastname
		name1 = name {
			fname: firstName, 
			lname: lastName,
		}
		// add each struct to the slice
		names = append(names, name1)		
	}
	
	fmt.Println(names)	
}