package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		displayUsage()
		return
	}

	task := strings.ToLower(args[0])
	switch task {
	case "uuid":
		fmt.Println(generateUUIDv7())
	default:
		displayUsage()
	}
}

func displayUsage() {
	fmt.Println("Usage: gen.exe [task] (params...)")
	fmt.Println()
	fmt.Println("gen uuid")
}
