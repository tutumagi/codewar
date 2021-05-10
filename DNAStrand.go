package main

func DNAStrand(dna string) string {
	mm := map[rune]rune{
		'T': 'A',
		'A': 'T',
		'C': 'G',
		'G': 'C',
	}
	ret := make([]rune, len(dna))
	for i, s := range dna {
		ret[i] = mm[s]
	}
	return string(ret)
}
