package main

import (
	"os"

	"timothyandrew.net/advent-2020/impl"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "1":
		impl.One()
	case "2":
		impl.Two()
	case "3":
		impl.Three()
	case "4":
		impl.Four()
	case "5":
		impl.Five()
	case "6":
		impl.Six()
	case "7":
		impl.Seven()
	case "8":
		impl.Eight()
	}
}
