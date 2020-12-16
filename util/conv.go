package util

import "strconv"

func ToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic("Failed to convert string to int")
	}

	return i
}

func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic("Failed to convert string to int")
	}

	return i
}

func FromBinaryString(s string) int64 {
	n, err := strconv.ParseInt(string(s), 2, 64)

	if err != nil {
		panic("Invalid binary string")
	}

	return n
}
