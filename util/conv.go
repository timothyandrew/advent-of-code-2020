package util

import "strconv"

func ToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic("Failed to convert string to int")
	}

	return i
}
