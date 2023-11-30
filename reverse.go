package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"
)

var reverseFlag = flag.String("reverse", "example.txt", "read file from flag")

func reverse(args []string) {
	checkArgs()
	fonts := "standard.txt"
	const usage = "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>"
	if !strings.Contains(*reverseFlag, "--reverse=") && len(args) == 1 {
		fmt.Println(args[0])
		os.Exit(0)
	}
	if len(args) > 0 {
		fmt.Println(usage)
		return
	}

	input, err := ioutil.ReadFile("test/" + *reverseFlag)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
		return
	}

	matrix := strings.Split(string(input), "\n")
	spaces := findSpace(matrix)
	userInput := splitUserInput(matrix, spaces)
	userInputMap := userInputMapping(userInput)
	asciiGraphic := getASCIIgraphicFont(fonts)
	output := mapUserInputWithASCIIgraphicFont(userInputMap, asciiGraphic)
	fmt.Println(output)
}

func checkArgs() {
	if len(os.Args) < 2 || (strings.Contains(os.Args[1], "--") && !strings.Contains(os.Args[1], "=")) {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		os.Exit(0)
	}
}

func findSpace(matrix []string) []int {
	var emptyColumns []int
	for column := 0; column < len(matrix[0]); column++ {
		count := 0
		for row := 0; row < len(matrix)-1; row++ {
			if matrix[row][column] == ' ' {
				count++
			} else {
				count = 0
				break
			}
			if count == len(matrix)-1 {
				emptyColumns = append(emptyColumns, column)
				count = 0
			}
		}
	}

	count := 5
	var indexToRem []int
	for i := range emptyColumns {
		if count == 0 {
			count = 5
			continue
		}
		if i > 0 && emptyColumns[i] == emptyColumns[i-1]+1 {
			indexToRem = append(indexToRem, i)
			count--
		}
	}
	for i := len(indexToRem) - 1; i >= 0; i-- {
		emptyColumns = removeIndex(emptyColumns, indexToRem[i])
	}
	return emptyColumns
}

func removeIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func splitUserInput(matrix []string, emptyColumns []int) string {
	var result strings.Builder
	result.WriteString("\n")
	start := 0
	end := 0
	for _, column := range emptyColumns {
		if end < len(matrix[0]) {
			end = column
			for _, characters := range matrix {
				if len(characters) > 0 {
					columns := characters[start:end]
					result.WriteString(columns)
					result.WriteString(" ")
				}
				result.WriteString("\n")
			}
			start = end + 1
		}
	}
	return result.String()
}

func userInputMapping(result string) map[int][]string {
	strSlice := strings.Split(result, "\n")
	graphicInput := make(map[int][]string)
	j := 0
	for _, ch := range strSlice {
		if ch == "" {
			j++
		} else {
			graphicInput[j] = append(graphicInput[j], ch)
		}
	}
	return graphicInput
}

func getASCIIgraphicFont(fonts string) map[int][]string {
	readFile, err := ioutil.ReadFile(fonts)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
		return nil
	}
	slice := strings.Split(string(readFile), "\n")
	ascii := make(map[int][]string)
	i := 31
	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}
	return ascii
}

func mapUserInputWithASCIIgraphicFont(graphicInput, ascii map[int][]string) string {
	var keys []int
	for k := range graphicInput {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var output string
	var sliceOfBytes []byte
	for _, value := range keys {
		graphicValue := graphicInput[value]
		for asciiKey, asciiValue := range ascii {
			if reflect.DeepEqual(asciiValue, graphicValue) {
				sliceOfBytes = append(sliceOfBytes, byte(asciiKey))
			}
		}
		output = string(sliceOfBytes)
	}
	return output
}
