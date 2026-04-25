func removeElement(nums []int, val int) int {
	k := len(nums)
	r := 0
	for _, num := range nums {
		if num == val {
			k--
		} else {
			nums[r] = num
			r++
		}
	}
	return k
}

