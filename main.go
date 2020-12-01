package main

import (
	"os"

	"timothyandrew.net/advent-2020/impl"
)

func main() {
	args := os.Args[1:]

	if args[0] == "1" {
		impl.One()
	}
}
