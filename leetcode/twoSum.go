package leetcode

func twoSum(nums []int, target int) []int {
	lens := len(nums)
	if lens < 2 {
		return []int{0, 0}
	}
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func towSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	var result = make([]int, 0, 2)
	m := make(map[int]int, len(nums))
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
			break
		}
		m[k] = i
	}
	return result
}
