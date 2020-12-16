package impl

import (
	"fmt"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

func IntersectionRuleSet(m1 map[rule]bool, m2 map[rule]bool) map[rule]bool {
	result := make(map[rule]bool)

	for k, v := range m1 {
		if _, ok := m2[k]; ok {
			result[k] = v
		}
	}

	return result
}

type span struct {
	lower int
	upper int
}

func (s span) validate(i int) bool {
	return i >= s.lower && i <= s.upper
}

func NewSpan(s string) span {
	matches := strings.Split(s, "-")
	return span{
		lower: util.ToInt(matches[0]),
		upper: util.ToInt(matches[1]),
	}
}

type rule struct {
	field string
	lspan span
	rspan span
}

func NewRule(s string) rule {
	matches := strings.Split(s, ":")
	field := matches[0]

	matches = strings.Split(matches[1], "or")
	lspan := NewSpan(strings.Trim(matches[0], " "))
	rspan := NewSpan(strings.Trim(matches[1], " "))

	return rule{field: field, lspan: lspan, rspan: rspan}
}

type ticket struct {
	values []int
}

func NewTicket(s string) ticket {
	matches := strings.Split(s, ",")

	var values []int
	for _, match := range matches {
		values = append(values, util.ToInt(match))
	}

	return ticket{values}
}

func (t ticket) validate(rules []rule) (invalidFields []int) {
Values:
	for _, value := range t.values {
		for _, rule := range rules {
			if rule.lspan.validate(value) || rule.rspan.validate(value) {
				continue Values
			}
		}

		invalidFields = append(invalidFields, value)
	}

	return
}

func Sixteen() {
	lines := util.FileToLines("input/16.txt")

	var rules []rule
	i := 0
	for ; i < len(lines); i++ {
		if strings.Trim(lines[i], "\t\n ") == "" {
			break
		}

		rules = append(rules, NewRule(lines[i]))
	}

	i += 2
	me := NewTicket(lines[i])

	i += 3
	var nearby []ticket
	for ; i < len(lines); i++ {
		nearby = append(nearby, NewTicket(lines[i]))
	}

	sum := 0
	var validTickets []ticket
	for _, ticket := range nearby {
		invalidFields := ticket.validate(rules)

		if len(invalidFields) == 0 {
			validTickets = append(validTickets, ticket)
		}

		for _, invalidField := range invalidFields {
			sum += invalidField
		}
	}

	fmt.Println("PART 1:", sum)

	tickets := validTickets
	tickets = append(tickets, me)

	// Fields whose rule is known
	known := make(map[rule]int)

Outer:
	for len(known) < len(tickets[0].values) {
		for i := 0; i < len(tickets[0].values); i++ {
			// All rules this field matches
			var fieldSet map[rule]bool

			for _, ticket := range tickets {
				// All rules this ticket field matches
				ticketFieldSet := make(map[rule]bool)

				for _, r := range rules {
					if _, ok := known[r]; ok {
						continue
					}

					if r.lspan.validate(ticket.values[i]) || r.rspan.validate(ticket.values[i]) {
						ticketFieldSet[r] = true
					}
				}

				if fieldSet == nil {
					fieldSet = ticketFieldSet
				} else {
					fieldSet = IntersectionRuleSet(fieldSet, ticketFieldSet)
				}
			}

			if len(fieldSet) == 1 {
				// Ugh
				for rule := range fieldSet {
					known[rule] = i
				}
				continue Outer
			}
		}
	}

	mul := 1
	for rule, i := range known {
		if strings.HasPrefix(rule.field, "departure") {
			mul *= me.values[i]
		}
	}

	fmt.Println("PART 2:", mul)
}
