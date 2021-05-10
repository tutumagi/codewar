package main

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	sb := &strings.Builder{}

	sb.WriteString("(")
	for idx, n := range numbers {
		if idx <= 2 {
			sb.WriteString(fmt.Sprintf("%d", n))
			if idx == 2 {
				sb.WriteString(") ")
			}
		} else if idx <= 5 {
			sb.WriteString(fmt.Sprintf("%d", n))
			if idx == 5 {
				sb.WriteString("-")
			}
		} else {
			sb.WriteString(fmt.Sprintf("%d", n))
		}
	}

	return sb.String()
}
