package main

func findDuplicate(nums []int) int {
	hash := make(map[int]int)

	for _, val := range nums {
		if _, ok := hash[val]; !ok {
			hash[val] = val
		} else {
			return val
		}
	}

	return 0

}
