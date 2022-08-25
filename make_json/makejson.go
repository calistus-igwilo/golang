/*
Write a program which prompts the user to first enter a name, and
then enter an address. Your program should create a map and add
the name and address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var name string
var address string
var str string
func main () {

	reader := bufio.NewReader(os.Stdin)
	// read name
	fmt.Println("Enter your name, please :")
	name, _ := reader.ReadString('\n')
	// remove the delimeter from the string
	name = strings.TrimSuffix(name, "\n")

	// read address
	fmt.Println("Enter your address, please :")
	address, _ := reader.ReadString('\n')
	// remove the delimeter from the string
	address = strings.TrimSuffix(address, "\n")

	// make map
	var idMap map[string]string
	idMap = make(map[string]string)

	// assign values to map
	idMap["Name"] = name
	idMap["Address"] = address

	barr, _ := json.Marshal(idMap)

	fmt.Println(string(barr))

}