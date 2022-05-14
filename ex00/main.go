package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		os.Stderr.Write([]byte("Error\n"))
		os.Exit(1)
	}
	data, err := os.ReadFile(args[1])
	if err != nil {
		os.Stderr.Write([]byte("Error\n"))
		os.Exit(1)
	}
	strs := piscine.SplitWhiteSpaces(string(data))
	if !piscine.Is_valid_param(strs) {
		os.Stderr.Write([]byte("Error\n"))
		os.Exit(1)
	}
}

func getArgc(args []string) (l int) {
	for range args {
		l++
	}
	return
}
