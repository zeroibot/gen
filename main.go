package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const targetAll string = "all"

func main() {
	args := os.Args[1:]
	numArgs := len(args)
	if numArgs == 0 {
		displayUsage(targetAll)
		return
	}

	task := strings.ToLower(args[0])
	switch task {
	case "uuid":
		fmt.Println(generateUUIDv7())
	case "code":
		if numArgs < 3 {
			displayUsage("code")
			return
		}
		length, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid length")
			return
		}
		pattern := args[2]
		fmt.Println(generateCode(length, pattern))
	default:
		displayUsage(targetAll)
	}
}

func displayUsage(target string) {
	fmt.Println("Usage: gen.exe [task] (params...)")
	fmt.Println()
	if target == targetAll || target == "uuid" {
		fmt.Println("gen uuid")
	}
	if target == targetAll || target == "code" {
		fmt.Println("gen code [length] [pattern:aA1]")
	}
}
