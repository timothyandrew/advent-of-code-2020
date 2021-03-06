package util

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func FileToString(path string) string {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic("Failed to read file!")
	}

	// This assumes that the byte array contains UTF-8
	s := string(bytes)
	s = strings.Trim(s, " \n\t\r")
	return s
}

func FileToLines(path string) []string {
	s := FileToString(path)
	return strings.Split(s, "\n")
}

func ReadFileRegexMatches(path string, re regexp.Regexp) [][]string {
	line := FileToString(path)
	matches := re.FindAllStringSubmatch(line, -1)
	return matches
}

func ReadFileLinesRegexMatches(path string, re regexp.Regexp) [][]string {
	lines := FileToLines(path)

	var result [][]string
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		result = append(result, matches[0][1:])
	}

	return result
}

func ReadFileInts(path string) []int {
	strInts := FileToLines(path)

	var ints []int

	for _, strInt := range strInts {
		number, err := strconv.Atoi(strInt)

		if err != nil {
			panic("Found a non-integer!")
		}

		ints = append(ints, number)
	}

	return ints
}
