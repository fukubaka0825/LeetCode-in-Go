package _001_two_sum


/* [Go]  1.Brute Force O(n2) */
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++{
		for j := i + 1; j < len(nums); j++{
			if nums[j] == target - nums[i]{
				return []int{i,j}
			}
		}
	}
	return nil
}
/* [Go] 2. Two-pass Hash Table O(n)  point: using complement*/
func twoSum2(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		indexMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _,ok := indexMap[complement]; ok && indexMap[complement] != i{
			return []int{i,indexMap[complement]}
		}
	}
	return nil
}
/* [Go] 3. One-pass Hash Table O(n)  point: using complement*/
func twoSum3(nums []int, target int) []int{
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _,ok := indexMap[complement]; ok {
			return []int{i,indexMap[complement]}
		}
		indexMap[nums[i]] = i
	}
	return nil
}