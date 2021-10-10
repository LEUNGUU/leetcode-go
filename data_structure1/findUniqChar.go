package data_structure1

import "math"

func firstUniqChar(s string) int {
	hashTable := make(map[string][]int)
	for idx, char := range s {
		hashTable[string(char)] = append(hashTable[string(char)], idx)
	}
	minIdx := math.MaxInt64
	for _, v := range hashTable {
		if len(v) == 1 && v[0] < minIdx {
			minIdx = v[0]
		}
	}
	if minIdx == math.MaxInt64 {
		return -1
	} else {
		return minIdx
	}
}
