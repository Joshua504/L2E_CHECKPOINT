package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bannerReader()
	checkBannerExist()
	readBannerLine()
	batchReader()
}

func contentSpliter() {
	content, err := os.ReadFile("sample.txt")
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
