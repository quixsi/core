package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command>")
		fmt.Println("Commands:")
		fmt.Println("  version - Show version information")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println("CLI Version: 0.0.1")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}