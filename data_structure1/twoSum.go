package data_structure1

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, val := range nums {
		_, ok := mapping[val]
		if ok {
			rc := []int{mapping[val], i}
			return rc
		} else {
			mapping[target-val] = i
		}
	}
	return nil
}
