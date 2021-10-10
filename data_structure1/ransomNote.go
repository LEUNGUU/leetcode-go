package data_structure1

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magazineMap := make(map[string]int)
	for _, char := range magazine {
		magazineMap[string(char)]++
	}
	for _, char := range ransomNote {
		num, ok := magazineMap[string(char)]
		if !ok || num <= 0 {
			return false
		} else {
			magazineMap[string(char)] = num - 1
		}
	}
	return true
}
