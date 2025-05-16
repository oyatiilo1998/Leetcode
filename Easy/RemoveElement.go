func removeElement(nums []int, val int) int {
	var resultCount int

	for i := 0; i < len(nums) ; i++ {
		if nums[i] != val {
			 nums[resultCount] = nums[i] 
            resultCount++
		}
	}
    
	return resultCount
}
