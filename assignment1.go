package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the numbers seperated by space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := scanner.Text()


	fmt.Printf("%s \n", arr)
}