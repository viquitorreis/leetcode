// CONVERTER NÚMEROS ROMANOS ( STRING ) PARA INTEIROS

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		currStr := string(s[i])
		if i >= len(s)-1 {
			count += roman[currStr]
		} else {
			if roman[currStr] < roman[string(s[i+1])] {
				count -= roman[currStr]
			} else {
				count += roman[currStr]
			}
		}
	}

	return count

}