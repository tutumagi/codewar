package main

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
