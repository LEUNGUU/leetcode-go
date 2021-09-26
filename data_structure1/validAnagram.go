package data_structure1

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[string]int)
	for _, char := range s {
		sMap[string(char)]++
	}
	for _, char := range t {
		num, ok := sMap[string(char)]
		if !ok || num <= 0 {
			return false
		} else {
			sMap[string(char)] = num - 1
		}
	}
	return true
}
