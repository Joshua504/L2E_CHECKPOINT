package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//bannerReader()
	//checkBannerExist()
	//readBannerLine()
	//batchReader()
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("file does not exist...")
	}

	charMap := contentSpliter(string(content))
	//fmt.Println(charMap)

	result := getCharacter("hi", charMap)
	for _, line := range printAsciiRow(result) {
		fmt.Println(line)
	}

	fmt.Println(splitLine("hello\\nworld"))
}

func splitLine(line string) []string {
	lines := strings.Split(line, "\\n")

	return lines
}

// row print
func printAsciiRow(charMap [][]string) []string {
	row := make([]string, 8)

	for _, block := range charMap {
		for i := 0; i < 8; i++ {
			row[i] += block[i]
		}
	}

	return row
}

// 2D storage for characters
func getCharacter(str string, charMap map[rune][]string) [][]string {
	words := [][]string{}

	for _, char := range str {
		words = append(words, charMap[char])
	}
	return words
}

func contentSpliter(content string) map[rune][]string {
	lines := strings.Split(string(content), "\n")
	characterMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		startLine := (i - 32) * 9
		characterMap[rune(i)] = lines[startLine+1 : startLine+9]
	}

	return characterMap
}

//File System (fs) API

func bannerReader() {
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("no such file or directory", err)
	}

	fmt.Println(string(content))
}

func checkBannerExist() {
	open, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error: file not found", err)
	}
	defer func() {
		if err := open.Close(); err != nil {
			fmt.Println("Error: can not close the file")
		}
	}()
}

func readBannerLine() {
	banner, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("no such file or directory", err)
	}
	defer func() {
		if err := banner.Close(); err != nil {
			fmt.Println("Error: can not close the file")
		}
	}()

	scanner := bufio.NewScanner(banner)
	var store []string

	for scanner.Scan() {
		line := scanner.Text()

		store = append(store, line)
	}

	fmt.Println(store)
}

func batchReader() {
	fileName := []string{"sample.txt", "test.txt"}

	for _, file := range fileName {
		banner, err := os.Open(file)
		if err != nil {
			fmt.Println("no such file or directory", err)
			continue
		}

		scanner := bufio.NewScanner(banner)
		count := 0

		for scanner.Scan() {
			scanner.Text()
			count++
		}

		banner.Close()
		line := fmt.Sprintf("%s contains %d lines\n", file, count)
		fmt.Print(line)
	}
}
