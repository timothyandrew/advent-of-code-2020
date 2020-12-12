package impl

import (
	"fmt"
	"math"
	"regexp"

	"timothyandrew.net/advent-2020/util"
)

type coordinate struct{ x, y int }

type ferry struct {
	direction string
	position  coordinate
}

func NewFerry() ferry {
	return ferry{direction: "E", position: coordinate{0, 0}}
}

func (f *ferry) turnLeft() {
	switch f.direction {
	case "N":
		f.direction = "W"
	case "S":
		f.direction = "E"
	case "E":
		f.direction = "N"
	case "W":
		f.direction = "S"
	default:
		panic("Invalid direction!")
	}
}

// Doesn't handle negative degrees
func (f *ferry) turnLeftDegrees(degrees int) {
	turns := int(math.Floor(float64(degrees) / 90))
	for i := 0; i < turns; i++ {
		f.turnLeft()
	}
}

func (f *ferry) moveForward(steps int) {
	switch f.direction {
	case "N":
		f.position.y += steps
	case "S":
		f.position.y -= steps
	case "E":
		f.position.x += steps
	case "W":
		f.position.x -= steps
	default:
		panic("Invalid direction!")
	}
}

func (f *ferry) takeAction(i instruction) {
	switch i.action {
	case "N":
		f.position.y += i.value
	case "S":
		f.position.y -= i.value
	case "E":
		f.position.x += i.value
	case "W":
		f.position.x -= i.value
	case "L":
		f.turnLeftDegrees(i.value)
	case "R":
		f.turnLeftDegrees(360 - i.value)
	case "F":
		f.moveForward(i.value)
	default:
		panic("Invalid instruction!")
	}
}

type waypointedFerry struct {
	ferry
	waypoint coordinate
}

func NewWaypointedFerry() waypointedFerry {
	return waypointedFerry{ferry: NewFerry(), waypoint: coordinate{10, 1}}
}

func (f *waypointedFerry) turnWaypointLeft() {
	x, y := f.waypoint.x, f.waypoint.y

	f.waypoint.x = -y
	f.waypoint.y = x
}

func (f *waypointedFerry) turnWaypointLeftDegrees(degrees int) {
	turns := int(math.Floor(float64(degrees) / 90))
	for i := 0; i < turns; i++ {
		f.turnWaypointLeft()
	}
}

func (f *waypointedFerry) moveForwardToWaypoint(steps int) {
	for i := 0; i < steps; i++ {
		f.position.x += f.waypoint.x
		f.position.y += f.waypoint.y
	}
}

func (f *waypointedFerry) takeAction(i instruction) {
	switch i.action {
	case "N":
		f.waypoint.y += i.value
	case "S":
		f.waypoint.y -= i.value
	case "E":
		f.waypoint.x += i.value
	case "W":
		f.waypoint.x -= i.value
	case "L":
		f.turnWaypointLeftDegrees(i.value)
	case "R":
		f.turnWaypointLeftDegrees(360 - i.value)
	case "F":
		f.moveForwardToWaypoint(i.value)
	default:
		panic("Invalid instruction!")
	}
}

type instruction struct {
	action string
	value  int
}

func ParseInstruction(s string) instruction {
	re := regexp.MustCompile("(\\w)(\\d+)")
	matches := re.FindStringSubmatch(s)
	return instruction{action: matches[1], value: util.ToInt(matches[2])}
}

func Twelve() {
	lines := util.FileToLines("input/12.txt")
	ferry := NewFerry()

	for _, line := range lines {
		i := ParseInstruction(line)
		ferry.takeAction(i)
	}

	fmt.Println("PART 1:", ferry)

	ferry2 := NewWaypointedFerry()

	for _, line := range lines {
		i := ParseInstruction(line)
		ferry2.takeAction(i)
	}

	fmt.Println("PART 2:", ferry2)
}
