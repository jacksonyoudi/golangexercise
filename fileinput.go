package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("method.go")
	if err != nil {
		fmt.Println("An error")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, erro := inputReader.ReadString('\n')
		fmt.Println(inputString)
		if erro == io.EOF {
			return
		}
	}
}
