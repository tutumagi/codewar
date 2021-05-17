package main

// https://www.codewars.com/kata/52774a314c2333f0a7000688/go
func ValidParentheses(parens string) bool {
	validator := make([]struct{}, 0, len(parens))
	for _, v := range parens {
		if v == '(' {
			validator = append(validator, struct{}{})
		}
		if v == ')' {
			if len(validator) <= 0 {
				return false
			}
			validator = validator[1:]
		}
	}
	return len(validator) == 0
}
