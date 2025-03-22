package main

import (
	"strings"
)

func handler(command string) int {
	if command[len(command)-1] == '\n' {
		command = command[:len(command)-1]
	}
	var spitCmd = strings.Split(command, " ")
	var mjrCmd = spitCmd[0]
	switch mjrCmd {
	case "exit":
		return handleexit(spitCmd)
	case "echo":
		handleecho(spitCmd)
	case "type":
		handletype(spitCmd)

	default:
		invalidCommand(command)
	}
	return EXIT_SUCCESS

}
