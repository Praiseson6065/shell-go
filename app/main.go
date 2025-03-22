package main

import (
	"fmt"
	"io"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {

		cmd, err := getCMD()
		if err != nil {
			if err == io.EOF {
				fmt.Println("exit")
				os.Exit(lastExitCode)
			}
		}
		lastExitCode = handler(cmd)
	}

}
