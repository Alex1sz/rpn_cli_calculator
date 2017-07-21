package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// NewReader returns a new Reader whose buffer has the default size
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple shell...")
	fmt.Println("--------------------------------")

	for {
		fmt.Print("-> ")

		inputStr, _ := reader.ReadString('\n')
		// remove trailing new line character
		inputStr = strings.Replace(inputStr, "\n", "", -1)

		if strings.Compare("howdy there", inputStr) == 0 {
			fmt.Println("...well hello")
		}
	}
}
