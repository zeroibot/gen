package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: gen.exe [task] (params...)")
		return
	}

	task := strings.ToLower(args[0])
	fmt.Println(task)
}
