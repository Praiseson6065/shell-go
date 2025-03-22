package main

import (
	"bufio"
	"fmt"
	"os"
)

func getCMD() (string, error) {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	return bufio.NewReader(os.Stdin).ReadString('\n')

}

func invalidCommand(command string) {
	fmt.Printf("%s: command not found\n", command)
}
