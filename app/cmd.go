package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

func handletype(args []string) (code int) {
	if slices.Contains(shell_builtins, args[1]) {
		fmt.Printf("%s is a shell builtin\n", args[1])
	} else if path, err := exec.LookPath(args[1]); err == nil {
		fmt.Printf("%s is %s\n", args[1], path)
	} else {
		fmt.Printf("%s: not found\n", args[1])
	}
	return EXIT_SUCCESS
}

func handleecho(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func handleexit(args []string) int {
	if len(args) > 1 {
		exitCode, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid exit code")
			return 1
		}
		os.Exit(exitCode)
	}
	os.Exit(0)
	return 0
}
