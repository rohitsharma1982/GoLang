package strand

var transcriptionMap = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

func ToRNA(dna string) string {
	rna := ""

	if len(dna) != 0 {
		for _, character := range dna {
			rna += transcriptionMap[character]
		}
	}
	return rna
}
