package main

func SplitStrings(str string) []string {
	var ss = []rune{'_', '_'}
	result := make([]string, 0, len(str)/2+1)
	for i, v := range str {
		if i%2 == 0 {
			if i != 0 {
				result = append(result, string(ss))
			}
			ss = []rune{v, '_'}
		} else {
			ss[1] = v
		}
	}
	result = append(result, string(ss))

	return result
}
