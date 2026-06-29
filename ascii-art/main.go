package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: not enough Arguments___ Usage: go run . <text>")
	}

	text := os.Args[1]
	if text == "" {
		return
	}

	banner, err := os.ReadFile("thinkertoy.txt")
	errorHandling(err)

	toString := string(banner)
	toString = strings.ReplaceAll(toString, "\r\n", "\n")

	splitBanner := strings.Split(toString, "\n")
	charMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		startLine := (i - 32) * 9
		charLine := splitBanner[startLine+1 : startLine+9]
		charMap[rune(i)] = charLine
	}
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				fmt.Println()
			}
			continue
		}

		row := make([]string, 8)

		for _, char := range line {
			for i := 0; i < 8; i++ {
				row[i] += charMap[char][i]
			}
		}

		for _, r := range row {
			fmt.Println(r)
		}
	}
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
