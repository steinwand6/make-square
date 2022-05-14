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
	for _, str := range strs {
		println(str)
		println(len(str))
		if len(str) != 16 || CountIf(str, '#') != 4 || CountIf(str, '.') != 12 {
			os.Stderr.Write([]byte("Error\n"))
			os.Exit(1)
		}
		if !is_valid_tetrimino(str) {
			os.Stderr.Write([]byte("Error\n"))
			os.Exit(1)
		}
	}
}

func is_valid_tetrimino(m string) bool {
	found := 0
	for i, r := range m {
		if r == '.' {
			continue
		}
		if i > 1 && i%4 != 0 && m[i-1] == '#' {
			found++
		}
		if i < 15 && m[i+1] == '#' {
			found++
		}
		if i >= 4 && m[i-4] == '#' {
			found++
		}
		if i < 12 && m[i+4] == '#' {
			found++
		}
	}
	return found == 6 || found == 8
}

func CountIf(s string, to_find rune) (result int) {
	for _, r := range s {
		if r == to_find {
			result++
		}
	}
	return
}

func getArgc(args []string) (l int) {
	for range args {
		l++
	}
	return
}
