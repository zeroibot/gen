package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	targetAll    string = "all"
	taskUUID     string = "uuid"
	taskCode     string = "code"
	taskPassword string = "password"
)

func main() {
	args := os.Args[1:]
	numArgs := len(args)
	if numArgs == 0 {
		displayUsage(targetAll)
		return
	}

	task := strings.ToLower(args[0])
	switch task {
	case taskUUID:
		fmt.Println(generateUUIDv7())
	case taskCode:
		if numArgs < 3 {
			displayUsage(taskCode)
			return
		}
		length, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid length")
			return
		}
		pattern := args[2]
		fmt.Println(generateCode(length, pattern))
	case taskPassword:
		if numArgs < 2 {
			displayUsage(taskPassword)
			return
		}
		rawPassword := args[1]
		password, err := hashPassword(rawPassword)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(password)
	default:
		displayUsage(targetAll)
	}
}

func displayUsage(target string) {
	fmt.Println("Usage: gen.exe [task] (params...)")
	fmt.Println()
	if target == targetAll || target == taskUUID {
		fmt.Println("gen uuid")
	}
	if target == targetAll || target == taskCode {
		fmt.Println("gen code [length] [pattern:aA1]")
	}
	if target == targetAll || target == taskPassword {
		fmt.Println("gen password [rawPassword]")
	}
}
