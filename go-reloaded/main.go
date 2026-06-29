package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <sampe.txt> <result.txt>")
		return
	}

	baseFile := os.Args[1]

	fileByte, err := os.ReadFile(baseFile)
	if err != nil {
		fmt.Println("Error: File does not exist in the directory")
		return
	}

	toStr := string(fileByte)
	splitStr := strings.Fields(toStr)

	result := []string{}

	skipNext := false

	for i, marker := range splitStr {
		if skipNext {
			skipNext = false
			continue
		}

		switch {
		case marker == "(hex)":
			toHex := convertToDecimal(splitStr, marker, "(hex)", 16, i)
			result[len(result)-1] = toHex
		case marker == "(bin)":
			toBin := convertToDecimal(splitStr, marker, "(bin)", 2, i)
			result[len(result)-1] = toBin
		case marker == "(up)":
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
		case marker == "(low)":
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
		case marker == "(cap)":
			result[len(result)-1] = capitalize(result[len(result)-1])
		case strings.HasPrefix(marker, "(up,") && strings.HasSuffix(splitStr[i+1], ")"):
			getNum := extractNumber(splitStr[i+1])
			result = applyChange(result, getNum, strings.ToUpper)
			skipNext = true
		case strings.HasPrefix(marker, "(low,") && strings.HasSuffix(splitStr[i+1], ")"):
			getNum := extractNumber(splitStr[i+1])
			result = applyChange(result, getNum, strings.ToLower)
			skipNext = true
		case strings.HasPrefix(marker, "(cap,") && strings.HasSuffix(splitStr[i+1], ")"):
			getNum := extractNumber(splitStr[i+1])
			result = applyChange(result, getNum, capitalize)
			skipNext = true
			fmt.Println(result)

		default:
			result = append(result, marker)
		}
	}

}

func convertToDecimal(array []string, text, con string, base, i int) string {
	if text == con {
		prev := array[i-1]
		convert, err := strconv.ParseInt(prev, base, 64)
		errorHandling(err)
		array[i-1] = strconv.FormatInt(convert, 10)
		return array[i-1]
	}
	return ""
}

func capitalize(str string) string {
	firstChar := strings.ToUpper(string(str[0]))
	remaining := strings.ToLower(str[1:])
	combine := firstChar + remaining

	return combine
}

func extractNumber(str string) int {
	trimBraces := strings.TrimRight(str, ")")
	toInt, err := strconv.Atoi(trimBraces)
	errorHandling(err)
	return toInt
}

func applyChange(result []string, digit int, lib func(string) string) []string {
	if digit > len(result) {
		digit = len(result)
	}
	for i := len(result) - digit; i < len(result); i++ {
		convert := lib(result[i])
		result[i] = convert
	}
	return result
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
