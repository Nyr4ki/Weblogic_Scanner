package main

import (
	"Weblogic_Scanner/poc/Console"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "%s [IP] [PORT]\n", args[0])
		return
	}
	url := args[1]
	port := args[2]

	Console.Run(url, port)
}
