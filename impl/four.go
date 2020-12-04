package impl

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

type field struct {
	name  string
	value string
}

type document struct {
	fields map[string]field
}

func buildInput() []document {
	file := util.FileToString("input/4.txt")
	lines := strings.Split(file, "\n\n")

	var documents []document

	for _, line := range lines {
		var (
			fields = make(map[string]field)
			re     = regexp.MustCompile("(\\S+):(\\S+)")
		)

		line = strings.ReplaceAll(line, "\n", " ")
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fields[match[1]] = field{name: match[1], value: match[2]}
		}

		documents = append(documents, document{fields})
	}

	return documents
}

func isValid(d document) bool {
	requiredFieldNames := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, f := range requiredFieldNames {
		if _, ok := d.fields[f]; !ok {
			return false
		}
	}

	if byr, ok := strconv.Atoi(d.fields["byr"].value); ok != nil || byr < 1920 || byr > 2002 {
		return false
	}

	if iyr, ok := strconv.Atoi(d.fields["iyr"].value); ok != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr, ok := strconv.Atoi(d.fields["eyr"].value); ok != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt
	re := regexp.MustCompile("^(\\d+)(cm|in)$")
	matches := re.FindStringSubmatch(d.fields["hgt"].value)
	if matches == nil {
		return false
	}
	nStr, unit := matches[1], matches[2]
	n, ok := strconv.Atoi(nStr)
	if ok != nil {
		return false
	}
	if unit == "cm" && (n < 150 || n > 193) {
		return false
	}
	if unit == "in" && (n < 59 || n > 76) {
		return false
	}

	// hcl
	re = regexp.MustCompile("^#[0-9a-f]{6}$")
	matches = re.FindStringSubmatch(d.fields["hcl"].value)
	if matches == nil {
		return false
	}

	// 🙄
	if ecl := d.fields["ecl"].value; ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return false
	}

	// pid
	re = regexp.MustCompile("^\\d{9}$")
	matches = re.FindStringSubmatch(d.fields["pid"].value)
	if matches == nil {
		return false
	}

	return true
}

func Four() {
	documents := buildInput()

	validDocuments := 0
	for _, d := range documents {
		if isValid(d) {
			validDocuments++
		}
	}

	fmt.Println("PART 1: Valid docs: ", validDocuments)
}
