package arrays101

func checkIfExist(arr []int) bool {
	hashTable := make(map[int]bool)
	for _, val := range arr {
		_, ok := hashTable[val]
		if ok {
			return true
		}
		hashTable[val*2] = true
		if val%2 == 0 {
			hashTable[val/2] = true
		}
	}
	return false
}
