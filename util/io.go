package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileInts(path string) []int {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic("Failed to read file!")
	}

	// This assumes that the byte array contains UTF-8
	s := string(bytes)
	s = strings.Trim(s, " \n\t\r")

	strInts := strings.Split(s, "\n")

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
