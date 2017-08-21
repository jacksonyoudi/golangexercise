package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter someting:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("there were errors reading, exiting program.")
		return
	}

	fmt.Printf("You name is %s", input)

	switch input {
	case "youdi\n":
		fmt.Println("Welcome youdi!")
	case "liang\r\n":
		fmt.Println("Welcome liang!")
	case "changyou\n":
		fmt.Println("Welcome changyou!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "youdi\n":
		fallthrough
	case "liang\n":
		fallthrough
	case "changyou\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "youdi\n", "liang\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

}
