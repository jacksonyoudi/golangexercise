package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "q\n" { // Windows, on Linux it is "S\n"
			os.Exit(0)
		}
		Cal(input)
	}
}

func Cal(input string) {
	a := strings.Fields(input)
	if len(a) != 3 {
		fmt.Println("error ,format is error")
		return
	}

	s1 := a[0]
	s2 := a[2]
	r1 := strconv.Atoi(s1)
	r2 := a[1]
	r3 := strconv.Atoi(s2)

	switch a[1] {
	case "+":
		fmt.Printf("The result of %d %s %d = %d\n", r1, r2, r3, r1+r3)
	case "-":
		fmt.Printf("The result of %d %s %d = %d\n", r1, r2, r3, r1-r3)
	case "*":
		fmt.Printf("The result of %d %s %d = %d\n", r1, r2, r3, r1*r3)
	case "/":
		fmt.Printf("The result of %d %s %d = %d\n", r1, r2, r3, r1/r3)
	default:
		fmt.Println("No valid input")
	}
}
