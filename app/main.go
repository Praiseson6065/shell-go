package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func invalidCommand(command string) {
	fmt.Printf("%s: command not found\n", command)
}

func handler(command string) {
	var args = strings.Split(command, " ")
	switch args[0] {
	case "exit":
		if len(args) > 1 {
			exitCode, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid exit code")
				return
			}
			os.Exit(exitCode)
		}
		return
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "type":
		if args[1]=="echo" || args[1]=="type" || args[1]=="exit" {
			fmt.Printf("%s is a shell builtin\n", args[1])
		}else {
			fmt.Printf("%s: not found\n", args[1])
		}
		
	default:
		invalidCommand(command)
	}

}

func main() {
	for {

		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		handler(cmd)
	}

}
