package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	m := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	m["name"] = strings.TrimSpace(name)
	fmt.Print("Enter Address: ")
	address, _ := reader.ReadString('\n')
	m["address"] = strings.TrimSpace(address)
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}